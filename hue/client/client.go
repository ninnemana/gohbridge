package client

import (
	"cloud.google.com/go/trace"
	"github.com/ninnemana/gohbridge/hue"
)

type client struct {
	trace *trace.Client
}

func New() (hue.Client, error) {
	return &client{}, nil
}
