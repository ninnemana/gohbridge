package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"cloud.google.com/go/trace"
	"github.com/ninnemana/gohbridge/services/bridge"
	"github.com/ninnemana/gohbridge/services/lights"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	tc, err := trace.NewClient(ctx, "ninneman-org")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(
		":50051",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(tc.GRPCClientInterceptor()),
	)
	if err != nil {
		log.Fatalf("did not connect RPC client: %v", err)
		return
	}
	defer conn.Close()

	bridgeConn := bridge.NewServiceClient(conn)

	brCall, err := bridgeConn.Discover(ctx, &bridge.DiscoverParams{
		Method: "remote",
	})
	if err != nil {
		log.Fatalf("failed to make gRPC discovery connection: %v", err)
		return
	}

	ipChan := make(chan interface{})
	go func() {
		for {
			br, err := brCall.Recv()
			switch err {
			case nil:
				ipChan <- br.GetInternalIPAddress()
			case io.EOF:
			default:
				ipChan <- err
			}
		}
	}()

	var ip string
	select {
	case result := <-ipChan:
		switch result.(type) {
		case error:
			log.Fatalf("failed to make discovery call to RPC service: %v", result)
		case string:
			ip = result.(string)
		}
	case <-time.After(time.Second * 5):
		log.Fatalf("timed out waiting for bridge discovery")
	}

	user := os.Getenv("HUE_USER")
	host := fmt.Sprintf("http://%s", ip)

	_, err = bridgeConn.GetBridgeState(ctx, &bridge.ConfigParams{
		User: user,
		Host: host,
	})
	if err != nil {
		panic(err)
	}

	lightConn := light.NewServiceClient(conn)

	lightCall, err := lightConn.All(ctx, &light.ListParams{
		User: user,
		Host: host,
	})

	var bedroom *light.Light
	for {
		l, err := lightCall.Recv()
		switch err {
		case nil:
			l, err := lightConn.Get(ctx, &light.GetParams{
				ID:   l.ID,
				User: user,
				Host: host,
			})
			if err != nil {
				fmt.Println("get: ", err)
				log.Fatal(err)
				return
			}

			if strings.TrimSpace(strings.ToLower(l.GetName())) == "bedroom lamp" {
				bedroom = l
			}
		case io.EOF:
			fmt.Println(bedroom.GetModelid())
			return
		default:
			fmt.Println("default: ", err.Error())
			return
		}
	}
}
