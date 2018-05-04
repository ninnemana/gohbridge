package service

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/trace"
	light "github.com/ninnemana/gohbridge/hue/lights"
	"github.com/pkg/errors"
)

type Service struct{}

func (s *Service) All(params *light.ListParams, server light.Service_AllServer) error {
	span := trace.FromContext(server.Context())
	child := span.NewChild("hue.lights.all")
	defer child.Finish()

	client := http.Client{
		Timeout: time.Second * 5,
	}

	path := fmt.Sprintf("%s/api/%s/lights", params.GetHost(), params.GetUser())
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

	for id, l := range lights {
		id, err := strconv.Atoi(id)
		if err != nil {
			return err
		}
		l.ID = int32(id)

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

	path := fmt.Sprintf("%s/api/%s/lights/new", params.GetHost(), params.GetUser())
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

	path := fmt.Sprintf("%s/api/%s/lights", params.GetHost(), params.GetUser())
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

	path := fmt.Sprintf("%s/api/%s/lights/%d", params.GetHost(), params.GetUser(), params.GetID())
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

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var l *light.Light
	if err := json.Unmarshal(data, &l); err != nil {
		return nil, errors.Errorf("failed to encode '%s' to Light: %v", data, err)
	}

	return l, nil
}

func (s *Service) SetState(ctx context.Context, params *light.SetStateParams) (*light.Light, error) {
	span := trace.FromContext(ctx)
	child := span.NewChild("hue.lights.set.state")
	defer child.Finish()

	client := http.Client{
		Timeout: time.Second * 5,
	}

	data, err := json.Marshal(params.GetUpdate())
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/api/%s/lights/%d/state", params.GetHost(), params.GetUser(), params.GetID())
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(data))
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

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res []interface{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	for _, r := range res {
		switch r.(type) {
		case map[string]interface{}:
			result := r.(map[string]interface{})
			switch {
			case result["error"] != nil:
				switch result["error"].(type) {
				case map[string]interface{}:
					e := result["error"].(map[string]interface{})
					switch e["description"].(type) {
					case string:
						return nil, errors.Errorf("failed to set state: %s", e["description"])
					}
				}

				return nil, errors.Errorf("state update failed")
			}
		default:
			return nil, errors.Errorf("failed to read state response")
		}
	}

	return s.Get(ctx, &light.GetParams{
		User: params.GetUser(),
		Host: params.GetHost(),
		ID:   params.GetID(),
	})
}
