package service

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"testing"

	"github.com/ninnemana/gohbridge/services/bridge"
	bridgeService "github.com/ninnemana/gohbridge/services/bridge/service"
	light "github.com/ninnemana/gohbridge/services/lights"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/trace"
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

	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})

	s := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))

	lightSvc, err := New()
	if err != nil {
		log.Fatal(err)
	}
	light.RegisterServiceServer(s, lightSvc)

	bridgeSvc, err := bridgeService.New()
	if err != nil {
		log.Fatal(err)
	}

	bridge.RegisterServiceServer(s, bridgeSvc)
	reflection.Register(s)

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to start RPC server: %s", err.Error())
		}
	}()

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithStatsHandler(&ocgrpc.ClientHandler{}))
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
	ctx := trace.NewContext(context.Background(), trace.NewSpan("test.lights.all", nil, trace.StartOptions{
		Sampler: trace.AlwaysSample(),
	}))
	for i := 0; i < 25; i++ {

		client, err := c.All(ctx, &light.ListParams{
			User: os.Getenv("HUE_USER"),
			Host: host,
		})
		if err != nil {
			t.Error(err)
			return
		}

		count := 0
		for {
			_, err = client.Recv()
			switch {
			case err == nil:
				count++
			case err == io.EOF:
				t.Logf("Light count '%d'\n", count)
				return
			default:
				t.Error(err)
				return
			}
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
