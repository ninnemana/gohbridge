package client

import "github.com/ninnemana/gohbridge/hue"

func (c *client) AllScenes() ([]interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) GetScene(string) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) CreateScene(interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) SetScene(string, interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) DeleteScene(string) error {
	return hue.ErrNotImplemented
}
