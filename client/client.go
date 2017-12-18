package main

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/trace"
	bridgepb "github.com/ninnemana/gohbridge/hue/bridge/service"
	"google.golang.org/grpc"
)

func main() {
	ctx := context.Background()
	tc, err := trace.NewClient(ctx, "ninneman-org")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := grpc.Dial(":50051", grpc.WithInsecure(), grpc.WithUnaryInterceptor(tc.GRPCClientInterceptor()))
	if err != nil {
		log.Fatalf("did not connect RPC client: %v", err)
		return
	}
	defer conn.Close()

	c := bridgepb.NewHueClient(conn)
	state, err := c.GetBridgeState(ctx, &bridgepb.Bridge{})
	if err != nil {
		panic(err)
	}
	fmt.Println(state)
}
