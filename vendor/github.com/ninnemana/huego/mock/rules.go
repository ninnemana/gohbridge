package client

import "github.com/ninnemana/huego"

func (c *client) AllRules() ([]interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) GetRule(string) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) CreateRule(interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) UpdateRule(string, interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) DeleteRule(string) error {
	return hue.ErrNotImplemented
}
