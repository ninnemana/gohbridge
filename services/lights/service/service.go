package service

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/trace"
	light "github.com/ninnemana/gohbridge/hue/lights"
	jsoniter "github.com/ninnemana/json-iterator"
	"github.com/pkg/errors"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// Service implements lights.Interactor
type Service struct{}

// All retrieves a streamed list of light records retrieved from the Hue REST API.
func (s *Service) All(params *light.ListParams, server light.Service_AllServer) error {
	child := trace.FromContext(server.Context()).NewChild("hue.lights.all")
	child.FinishWait()

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

// New polls the Hue Bridge for lights that are available to be setup.
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

// Search initializes a light search event.
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

// Get retrieves information about a specific light.
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

	l.ID = params.GetID()

	return l, nil
}

// SetState updates the current state of the defined light.
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

// Toggle changes the light from on -> off or off -> on depending on the current state of the light.
func (s *Service) Toggle(ctx context.Context, params *light.ToggleParams) (*light.Light, error) {
	span := trace.FromContext(ctx)
	child := span.NewChild("hue.lights.toggle")
	defer child.Finish()

	existing, err := s.Get(ctx, &light.GetParams{
		Host: params.GetHost(),
		ID:   params.GetID(),
		User: params.GetUser(),
	})
	if err != nil {
		return nil, err
	}

	alreadyOn := false
	switch {
	case existing.GetState() == nil:
	case existing.GetState().GetOn():
		alreadyOn = true
	case !existing.GetState().GetOn():
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	data := fmt.Sprintf(`{"on": %t}`, !alreadyOn)

	path := fmt.Sprintf("%s/api/%s/lights/%d/state", params.GetHost(), params.GetUser(), params.GetID())
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer([]byte(data)))
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

	var res []interface{}
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	for _, r := range res {
		switch r.(type) {
		case map[string]interface{}:
			result := r.(map[string]interface{})
			switch result["error"].(type) {
			case nil:
			case map[string]interface{}:
				e := result["error"].(map[string]interface{})
				switch e["description"].(type) {
				case string:
					return nil, errors.Errorf("failed to set state: %s", e["description"])
				}
				return nil, errors.Errorf("state update failed")
			}
		default:
			return nil, errors.Errorf("failed to read state response of '%T'", r)
		}
	}

	return s.Get(ctx, &light.GetParams{
		User: params.GetUser(),
		ID:   params.GetID(),
		Host: params.GetHost(),
	})
}
