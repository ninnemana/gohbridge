package client

import "github.com/ninnemana/huego"

func (c *client) AllSensors() ([]interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) CreateSensor(interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) SearchSensors() error {
	return hue.ErrNotImplemented
}

func (c *client) NewSensors() ([]interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) GetSensor(string) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) SetSensor(string, interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) RenameSensor(string, string) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) DeleteSensor(string) error {
	return hue.ErrNotImplemented
}
