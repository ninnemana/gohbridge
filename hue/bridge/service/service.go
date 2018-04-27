package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"cloud.google.com/go/trace"
	"github.com/ninnemana/gohbridge/hue/bridge"
	"github.com/pkg/errors"
	context "golang.org/x/net/context"
)

type Service struct {
	clients map[string]string
}

func (s Service) Discover(params *bridge.DiscoverParams, serv bridge.Service_DiscoverServer) error {
	span := trace.FromContext(serv.Context())
	child := span.NewChild("hue.discover")
	defer child.Finish()

	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(http.MethodGet, "https://www.meethue.com/api/nupnp", nil)
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

	bridges := []bridge.Bridge{}
	err = json.NewDecoder(resp.Body).Decode(&bridges)
	if err != nil {
		return err
	}

	for _, br := range bridges {
		err = serv.Send(&br)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s Service) GetBridgeState(ctx context.Context, params *bridge.ConfigParams) (*bridge.BridgeState, error) {
	span := trace.FromContext(ctx)
	child := span.NewChild("hue.bridge_state")
	defer child.Finish()

	client := http.Client{
		Timeout: time.Second * 5,
	}

	path := fmt.Sprintf("%s/api/%s", params.GetHost(), params.GetUser())
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

	bs := bridge.BridgeState{}
	err = json.NewDecoder(resp.Body).Decode(&bs)
	if err != nil {
		return nil, err
	}

	return &bs, nil
}

func (s Service) GetConfig(ctx context.Context, params *bridge.ConfigParams) (*bridge.BridgeConfig, error) {
	span := trace.FromContext(ctx)
	child := span.NewChild("hue.full_config")
	defer child.Finish()

	client := http.Client{
		Timeout: time.Second * 5,
	}

	path := fmt.Sprintf("%s/api/%s/config", params.GetHost(), params.GetUser())
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

	conf := bridge.BridgeConfig{}
	err = json.NewDecoder(resp.Body).Decode(&conf)
	if err != nil {
		return nil, err
	}

	return &conf, nil
}
