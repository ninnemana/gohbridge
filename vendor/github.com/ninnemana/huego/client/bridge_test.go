package client

import (
	"context"
	"log"
	"testing"

	hue "github.com/ninnemana/huego"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/trace"
)

var (
	cl hue.Client
)

func TestMain(m *testing.M) {
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: "ninneman-org",
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

	cl, err = New()
	if err != nil {
		log.Fatalf("failed to create new hue.Client: %v", err)
	}

	m.Run()
}

func TestAllBridges(t *testing.T) {

	_, span := trace.StartSpan(context.Background(), "test.hue.http")
	defer span.End()

	ctx := trace.NewContext(context.Background(), span)

	results, err := cl.AllBridges(ctx, &hue.AllBridgeParams{
		Method: "remote",
	})
	if err != nil {
		t.Errorf("failed to get all briges: %v", err)
		return
	}

	for _, res := range results {
		bridge, ok := res.(map[string]interface{})
		if !ok {
			t.Errorf("expected 'map[string]interface{}' got '%T'", res)
			continue
		}

		t.Log(bridge["internalipaddress"])
	}
}
