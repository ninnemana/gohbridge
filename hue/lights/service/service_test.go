package service

import (
	"io"
	"log"
	"net"
	"os"
	"testing"

	"cloud.google.com/go/trace"
	light "github.com/ninnemana/gohbridge/hue/lights"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	// port    = "50051"
	address = "localhost:50051"
)

var (
	c light.ServiceClient
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

	light.RegisterServiceServer(s, &Service{})
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

	c = light.NewServiceClient(conn)

	m.Run()
}

func TestAll(t *testing.T) {
	client, err := c.All(context.Background(), &light.ListParams{
		User: os.Getenv("HUE_USER"),
	})
	if err != nil {
		t.Error(err)
		return
	}

	for {
		msg, err := client.Recv()
		switch {
		case err == nil:
		case err == io.EOF:
			return
		default:
			t.Error(err)
			return
		}

		t.Log(msg)

	}
}

// func TestGetBridgeState(t *testing.T) {
// 	state, err := c.GetBridgeState(context.Background(), &bridge.ConfigParams{
// 		User: os.Getenv("HUE_USER"),
// 	})
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	t.Log(state)
// }
