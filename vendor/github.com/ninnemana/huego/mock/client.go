package client

import "cloud.google.com/go/trace"

type client struct {
	trace *trace.Client
}

func New() (*client, error) {
	return &client{
		// trace: tr,
	}, nil
}
