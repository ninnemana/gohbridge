package client

import (
	"context"

	"github.com/ninnemana/huego"
)

func (c *client) AllLights(ctx context.Context) ([]interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) NewLights(ctx context.Context) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) SearchLights(ctx context.Context, deviceIDs []string) error {
	return hue.ErrNotImplemented
}

func (c *client) GetLight(ctx context.Context, id int) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) RenameLight(string, string) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) LightState(ctx context.Context, id int, state interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) Toggle(ctx context.Context, id int) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) DeleteLight(string) error {
	return hue.ErrNotImplemented
}
