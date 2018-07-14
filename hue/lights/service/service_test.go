package service

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"cloud.google.com/go/trace"
	"github.com/ninnemana/gohbridge/hue/bridge"
	bridgeService "github.com/ninnemana/gohbridge/hue/bridge/service"
	light "github.com/ninnemana/gohbridge/hue/lights"
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	address = "localhost:50052"
)

var (
	c    light.ServiceClient
	host string
)

func TestMain(m *testing.M) {

	lis, err := net.Listen("tcp", ":50052")
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
	bridge.RegisterServiceServer(s, &bridgeService.Service{})
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

	client, err := bridge.NewServiceClient(conn).Discover(context.Background(), &bridge.DiscoverParams{
		Method: "remote",
	})
	if err != nil {
		log.Fatalf("failed to create bridge client: %v", err)
		return
	}

	var br *bridge.Bridge
	for br == nil {
		br, err = client.Recv()
		if err != nil {
			log.Fatalf("failed to receive bridge: %v", err)
			return
		}
	}

	host = fmt.Sprintf("http://%s", br.GetInternalIPAddress())

	m.Run()
}

func TestAll(t *testing.T) {

	client, err := c.All(context.Background(), &light.ListParams{
		User: os.Getenv("HUE_USER"),
		Host: host,
	})
	if err != nil {
		t.Error(err)
		return
	}

	for {
		_, err = client.Recv()
		switch {
		case err == nil:
		case err == io.EOF:
			return
		default:
			t.Error(err)
			return
		}
	}
}

func TestGet(t *testing.T) {
	ctx := context.Background()
	result, err := c.All(ctx, &light.ListParams{
		User: os.Getenv("HUE_USER"),
		Host: host,
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	var l *light.Light
	for l == nil {
		l, err = result.Recv()
		switch {
		case err == nil:
			break
		case err == io.EOF:
			t.Fatal("no lights available")
			return
		default:
			t.Error(err)
			return
		}
	}

	_, err = c.Get(context.Background(), &light.GetParams{
		User: os.Getenv("HUE_USER"),
		Host: host,
		ID:   l.GetID(),
	})
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSetState(t *testing.T) {
	ctx := context.Background()
	result, err := c.All(ctx, &light.ListParams{
		User: os.Getenv("HUE_USER"),
		Host: host,
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	var l *light.Light
	for l == nil {
		li, err := result.Recv()
		switch {
		case err == nil:
			if li.GetName() == "Desk Lamp" {
				l = li
				break
			}
		case err == io.EOF:
			t.Fatal("no lights available")
			return
		default:
			t.Error(err)
			return
		}
	}

	lg, err := c.Get(context.Background(), &light.GetParams{
		User: os.Getenv("HUE_USER"),
		Host: host,
		ID:   l.GetID(),
	})
	if err != nil {
		t.Error(err)
		return
	}

	lg, err = c.SetState(context.Background(), &light.SetStateParams{
		Update: &light.LightState{
			On:  true,
			Bri: 100.00,
		},
		User: os.Getenv("HUE_USER"),
		Host: host,
		ID:   l.GetID(),
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	if lg.GetState().GetBri() != 100 {
		t.Fatalf("failed to update brightness to '60' responded with '%f'", lg.GetState().GetBri())
		return
	}
}

func TestToggle(t *testing.T) {
	ctx := context.Background()

	_, err := c.Toggle(ctx, &light.ToggleParams{})
	if err == nil {
		t.Error("expected error not to be nil with empty toggle params")
	}

	res, err := c.All(ctx, &light.ListParams{
		User: os.Getenv("HUE_USER"),
		Host: host,
	})
	if err != nil {
		t.Fatal(err)
		return
	}

	var l *light.Light
	for l == nil {
		li, err := res.Recv()
		if err != nil {
			t.Error("failed to retrieve light")
			return
		}

		if strings.Contains(li.GetName(), "Desk") {
			l = li
		}
	}

	_, err = c.Toggle(ctx, &light.ToggleParams{
		Host: host,
		ID:   l.GetID(),
		User: os.Getenv("HUE_USER"),
	})
	if err != nil {
		t.Fatal(err)
		return
	}
}
