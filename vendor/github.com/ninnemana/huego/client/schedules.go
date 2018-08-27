package client

import "github.com/ninnemana/huego"

func (c *client) AllSchedules() ([]interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) CreateSchedule(interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) GetSchedule(string) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) SetSchedule(string, interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) DeleteSchedule(string) error {
	return hue.ErrNotImplemented
}
