package service

import (
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	"cloud.google.com/go/trace"
	hueClient "github.com/ninnemana/gohbridge/hue/client"
	"github.com/ninnemana/gohbridge/services/bridge"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	// port    = "50051"
	address = "localhost:50051"
)

var (
	c bridge.ServiceClient
)

func TestMain(m *testing.M) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen on: %s", address)
	}

	ctx := context.Background()
	tc, err := trace.NewClient(ctx, "ninneman-org")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(tc.GRPCServerInterceptor()))

	cl, err := hueClient.New()
	if err != nil {
		log.Fatal(err)
	}

	bridge.RegisterServiceServer(s, &Service{
		hue: cl,
	})
	reflection.Register(s)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to start RPC server: %s", err.Error())
		}
	}()

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(tc.GRPCClientInterceptor()))
	if err != nil {
		log.Fatalf("did not connect RPC client: %v", err)
		return
	}
	defer conn.Close()

	c = bridge.NewServiceClient(conn)

	m.Run()
}

func TestDiscover(t *testing.T) {
	client, err := c.Discover(context.Background(), &bridge.DiscoverParams{
		Method: "remote",
	})
	if err != nil {
		t.Error(err)
		return
	}

	for {
		msg, err := client.Recv()
		if err != nil {
			t.Error(err)
		}

		if msg.GetId() == "" {
			t.Error("should have returned an id")
		}
		if msg.GetInternalIPAddress() == "" {
			t.Error("should have returned an internal IP address")
		}

		return
	}
}

func TestGetBridgeState(t *testing.T) {
	client, err := c.Discover(context.Background(), &bridge.DiscoverParams{
		Method: "remote",
	})
	if err != nil {
		t.Error(err)
		return
	}

	var br *bridge.Bridge
	for br == nil {
		br, err = client.Recv()
		if err != nil {
			t.Error(err)
		}
	}

	_, err = c.GetBridgeState(context.Background(), &bridge.ConfigParams{
		User: os.Getenv("HUE_USER"),
		Host: fmt.Sprintf("http://%s", br.GetInternalIPAddress()),
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetConfig(t *testing.T) {
	client, err := c.Discover(context.Background(), &bridge.DiscoverParams{
		Method: "remote",
	})
	if err != nil {
		t.Error(err)
		return
	}

	var br *bridge.Bridge
	for br == nil {
		br, err = client.Recv()
		if err != nil {
			t.Error(err)
		}
	}

	_, err = c.GetConfig(context.Background(), &bridge.ConfigParams{
		User: os.Getenv("HUE_USER"),
		Host: fmt.Sprintf("http://%s", br.GetInternalIPAddress()),
	})
	if err != nil {
		t.Error(err)
		return
	}
}
