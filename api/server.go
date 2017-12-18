package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"cloud.google.com/go/trace"
	bridgepb "github.com/ninnemana/gohbridge/hue/bridge/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

func main() {
	ctx := context.Background()
	tc, err := trace.NewClient(ctx, "ninneman-org")
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on: %s", port)
	}

	grpc.EnableTracing = true

	s := grpc.NewServer(grpc.UnaryInterceptor(tc.GRPCServerInterceptor()))
	bridgepb.RegisterHueServer(s, &bridgepb.Service{})
	reflection.Register(s)

	fmt.Printf("starting RPC server on %s ...\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to start RPC server: %s", err.Error())
	}
}
