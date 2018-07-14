package logger

import (
	"context"

	"cloud.google.com/go/bigquery"
	"google.golang.org/grpc"
)

// Option define the configuration parameters for a logger interceptor.
type Option struct{}

// UnaryServerInterceptor statisfies middleware requirements for implementing
// a gRPC middleware service for logging incoming requests.
func UnaryServerInterceptor(bq *bigquery.Client, opts ...Option) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		return handler(ctx, req)
	}
}

// StreamServerInterceptor statisfies middleware requirements for implementing
// a gRPC middleware service for logging stream requests.
func StreamServerInterceptor(bq *bigquery.Client, opts ...Option) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		return handler(srv, stream)
	}
}
