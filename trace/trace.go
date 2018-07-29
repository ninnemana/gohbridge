package trace

import (
	"context"
	"fmt"

	gtrace "cloud.google.com/go/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const grpcMetadataKey = "grpc-trace-bin"

// Client implements a gRPC middleware interface
// for logging events into the google-cloud trace API.
type Client struct {
	trace *gtrace.Client
}

// NewClient initializes a new implementation of the Client.
func NewClient(ctx context.Context, projectID string) (*Client, error) {
	tc, err := gtrace.NewClient(ctx, projectID)
	if err != nil {
		return nil, err
	}

	return &Client{
		trace: tc,
	}, nil
}

// UnaryInterceptor returns a grpc.UnaryServerInterceptor that enables the tracing of the incoming
// gRPC calls. Incoming call's context can be used to extract the span on servers that enabled this option:
//
//	span := trace.FromContext(ctx)
//
// If the client is nil, then the interceptor just invokes the handler.
//
// The functionality in gRPC that this feature relies on is currently experimental.
func (c *Client) UnaryInterceptor() grpc.UnaryServerInterceptor {
	if c == nil {
		return func(ctx context.Context, req interface{}, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
			return handler(ctx, req)
		}
	}
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, _ := metadata.FromIncomingContext(ctx)
		var traceHeader string
		if header, ok := md[grpcMetadataKey]; ok {
			traceID, spanID, opts, ok := decode([]byte(header[0]))
			if ok {
				// TODO(jbd): Generate a span directly from string(traceID), spanID and opts.
				traceHeader = fmt.Sprintf("%x/%d;o=%d", traceID, spanID, opts)
			}
		}
		span := c.trace.SpanFromHeader(info.FullMethod, traceHeader)
		defer span.Finish()
		ctx = gtrace.NewContext(ctx, span)
		return handler(ctx, req)
	}
}

// StreamServerInterceptor implements the middleware requirements for the trace client
// to be implemented as stream request middleware.
func (c *Client) StreamServerInterceptor() grpc.StreamServerInterceptor {
	if c == nil {
		return func(
			srv interface{},
			stream grpc.ServerStream,
			info *grpc.StreamServerInfo,
			handler grpc.StreamHandler,
		) error {
			return handler(srv, stream)
		}
	}
	return func(
		srv interface{},
		stream grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		md, _ := metadata.FromIncomingContext(stream.Context())
		var traceHeader string
		if header, ok := md[grpcMetadataKey]; ok {
			traceID, spanID, opts, ok := decode([]byte(header[0]))
			if ok {
				// TODO(jbd): Generate a span directly from string(traceID), spanID and opts.
				traceHeader = fmt.Sprintf("%x/%d;o=%d", traceID, spanID, opts)
			}
		}
		stream.SetHeader(md)
		span := c.trace.SpanFromHeader(info.FullMethod, traceHeader)
		defer span.Finish()

		return handler(srv, stream)
	}
}
