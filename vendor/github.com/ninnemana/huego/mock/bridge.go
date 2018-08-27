package client

import (
	"context"

	"github.com/ninnemana/huego"
)

func (c *client) AllBridges(ctx context.Context, q interface{}) ([]interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) CreateUser(interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) GetConfig(ctx context.Context) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) ModifyConfig(interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) Unwhitelist(string) error {
	return hue.ErrNotImplemented
}

func (c *client) GetFullState(ctx context.Context) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}
