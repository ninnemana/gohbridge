package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/ninnemana/gohbridge/services/bridge"
	bridgeService "github.com/ninnemana/gohbridge/services/bridge/service"
	"github.com/ninnemana/gohbridge/services/lights"
	lightService "github.com/ninnemana/gohbridge/services/lights/service"

	huego "github.com/ninnemana/huego"
	hueSvc "github.com/ninnemana/huego/client"
	hueMock "github.com/ninnemana/huego/mock"

	"cloud.google.com/go/bigquery"
	"go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	port = ":50051"
)

func main() {
	ctx := context.Background()
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		MetricPrefix: "hue",
		ProjectID:    "ninneman-org",
		OnError: func(err error) {
			log.Fatal(err)
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer exporter.Flush()

	trace.RegisterExporter(exporter)
	defer trace.UnregisterExporter(exporter)
	view.RegisterExporter(exporter)
	defer view.UnregisterExporter(exporter)

	trace.ApplyConfig(trace.Config{
		DefaultSampler: trace.AlwaysSample(),
	})

	bq, err := bigquery.NewClient(ctx, "ninneman-org")
	if err != nil {
		log.Fatal(err)
	}
	_ = bq

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen on: %s", port)
	}

	grpc.EnableTracing = true

	s := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))
	// s := grpc.NewServer(
	// 	grpc.StreamInterceptor(
	// 		grpc_middleware.ChainStreamServer(
	// 			logger.StreamServerInterceptor(bq, logger.Option{}),
	// 		),
	// 	),
	// 	grpc.UnaryInterceptor(
	// 		grpc_middleware.ChainUnaryServer(
	// 			logger.UnaryServerInterceptor(bq, logger.Option{}),
	// 		),
	// 	),
	// )

	lg := grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr)
	grpclog.SetLoggerV2(lg)

	var hueClient huego.Client
	switch os.Getenv("HUE_MOCK") {
	case "":
		hueClient, err = hueSvc.New()
	default:
		hueClient, err = hueMock.New()
	}

	bridgeSvc, err := bridgeService.New(hueClient, lg)
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
