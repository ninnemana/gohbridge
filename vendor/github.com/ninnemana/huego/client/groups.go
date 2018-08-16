package client

import "github.com/ninnemana/huego"

func (c *client) AllGroups() ([]interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) CreateGroup(interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) GetGroup(string) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) SaveGroup(string, interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) SetGroupState(string, interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) DeleteGroup(string) error {
	return hue.ErrNotImplemented
}
