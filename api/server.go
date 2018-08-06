package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/ninnemana/gohbridge/hue"
	hueSvc "github.com/ninnemana/gohbridge/hue/client"
	hueMock "github.com/ninnemana/gohbridge/hue/mock"
	"github.com/ninnemana/gohbridge/logger"
	"github.com/ninnemana/gohbridge/services/bridge"
	bridgeService "github.com/ninnemana/gohbridge/services/bridge/service"
	"github.com/ninnemana/gohbridge/services/lights"
	lightService "github.com/ninnemana/gohbridge/services/lights/service"

	"cloud.google.com/go/bigquery"
	"cloud.google.com/go/trace"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
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

	var hueClient hue.Client
	switch os.Getenv("HUE_MOCK") {
	case "":
		hueClient, err = hueSvc.New()
	default:
		hueClient, err = hueMock.New()
	}

	bridgeSvc, err := bridgeService.New(hueClient, lg, tc)
	if err != nil {
		log.Fatalf("failed to start bridge service: %v", err)
	}

	lightSvc, err := lightService.New(hueClient)
	if err != nil {
		log.Fatalf("failed to start light service: %v", err)
	}

	bridge.RegisterServiceServer(s, bridgeSvc)
	light.RegisterServiceServer(s, lightSvc)

	// register bridge service
	// bridge.RegisterServiceServer(s, &bridgeSvc.Service{
	// 	Log:   lg,
	// 	Trace: tc,
	// })

	// register light service
	// light.RegisterServiceServer(s, &lightSvc.Service{
	// 	hue: client,
	// })

	grpclog.Infof("starting RPC server on %s ...\n", port)
	if err := s.Serve(lis); err != nil {
		grpclog.Fatalf("failed to start RPC server: %v", err)
	}
}
