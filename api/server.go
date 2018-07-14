package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/ninnemana/gohbridge/hue/bridge"
	bridgeSvc "github.com/ninnemana/gohbridge/hue/bridge/service"
	"github.com/ninnemana/gohbridge/hue/lights"
	lightSvc "github.com/ninnemana/gohbridge/hue/lights/service"
	"github.com/ninnemana/gohbridge/logger"
	"google.golang.org/grpc/grpclog"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/trace"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	ctx := context.Background()
	tc, err := trace.NewClient(ctx, "ninneman-org")

	bq, err := bigquery.NewClient(ctx, "ninneman-org")
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on: %s", port)
	}

	grpc.EnableTracing = true

	s := grpc.NewServer(
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				logger.StreamServerInterceptor(bq, logger.Option{}),
				// tc.StreamServerInterceptor(),
			),
		),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				logger.UnaryServerInterceptor(bq, logger.Option{}),
				tc.GRPCServerInterceptor(),
			),
		),
	)

	lg := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(lg)

	// register bridge service
	bridge.RegisterServiceServer(s, &bridgeSvc.Service{
		Log:   lg,
		Trace: tc,
	})

	// register light service
	light.RegisterServiceServer(s, &lightSvc.Service{})

	grpclog.Infof("starting RPC server on %s ...\n", port)
	if err := s.Serve(lis); err != nil {
		grpclog.Fatalf("failed to start RPC server: %v", err)
	}
}
