package client

import (
	"cloud.google.com/go/trace"
	"github.com/ninnemana/huego"
)

type client struct {
	trace *trace.Client
}

func New() (hue.Client, error) {
	return &client{}, nil
}
