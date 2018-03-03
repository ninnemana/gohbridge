package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"cloud.google.com/go/trace"
	light "github.com/ninnemana/gohbridge/hue/lights"
)

type Service struct{}

func (s *Service) All(params *light.ListParams, server light.Service_AllServer) error {
	span := trace.FromContext(server.Context())
	child := span.NewChild("hue.lights.all")
	defer child.Finish()

	client := http.Client{
		Timeout: time.Second * 5,
	}

	path := fmt.Sprintf("http://192.168.86.133/api/%s/lights", params.User)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(data))
	}

	lights := make(map[string]light.Light, 0)
	err = json.NewDecoder(resp.Body).Decode(&lights)
	if err != nil {
		return err
	}

	for _, l := range lights {
		if err = server.Send(&l); err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) New(ctx context.Context, params *light.NewParams) (*light.Scan, error) {
	span := trace.FromContext(ctx)
	child := span.NewChild("hue.lights.new")
	defer child.Finish()

	client := http.Client{
		Timeout: time.Second * 5,
	}

	path := fmt.Sprintf("http://192.168.86.133/api/%s/lights/new", params.User)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(data))
	}

	var scan *light.Scan
	err = json.NewDecoder(resp.Body).Decode(scan)
	if err != nil {
		return nil, err
	}

	return scan, nil
}

func (s *Service) Search(ctx context.Context, params *light.SearchParams) (*light.SearchResult, error) {
	span := trace.FromContext(ctx)
	child := span.NewChild("hue.lights.search")
	defer child.Finish()

	client := http.Client{
		Timeout: time.Second * 5,
	}

	path := fmt.Sprintf("http://192.168.86.133/api/%s/lights", params.User)
	req, err := http.NewRequest(http.MethodPost, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(data))
	}

	var result *light.SearchResult
	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *Service) Get(ctx context.Context, params *light.GetParams) (*light.Light, error) {
	span := trace.FromContext(ctx)
	child := span.NewChild("hue.lights.get")
	defer child.Finish()

	client := http.Client{
		Timeout: time.Second * 5,
	}

	path := fmt.Sprintf("http://192.168.86.133/api/%s/lights/%d", params.User, params.ID)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(data))
	}

	var l *light.Light
	err = json.NewDecoder(resp.Body).Decode(l)
	if err != nil {
		return nil, err
	}

	return l, nil
}
