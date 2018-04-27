package service

import (
	"fmt"
	"log"
	"net"
	"os"
	"testing"

	"cloud.google.com/go/trace"
	"github.com/ninnemana/gohbridge/hue/bridge"
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

	os.Setenv(DiscoverAddr, "http://www.meethue.com/api/nupnp")
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

	bridge.RegisterServiceServer(s, &Service{})
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

	os.Unsetenv(DiscoverAddr)
}

func TestDiscover(t *testing.T) {
	client, err := c.Discover(context.Background(), &bridge.DiscoverParams{})
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
	client, err := c.Discover(context.Background(), &bridge.DiscoverParams{})
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

	state, err := c.GetBridgeState(context.Background(), &bridge.ConfigParams{
		User: os.Getenv("HUE_USER"),
		Host: fmt.Sprintf("http://%s", br.GetInternalIPAddress()),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(state)
}

func TestGetConfig(t *testing.T) {
	client, err := c.Discover(context.Background(), &bridge.DiscoverParams{})
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

	config, err := c.GetConfig(context.Background(), &bridge.ConfigParams{
		User: os.Getenv("HUE_USER"),
		Host: fmt.Sprintf("http://%s", br.GetInternalIPAddress()),
	})
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(config)
}
