package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"cloud.google.com/go/trace"
	upnp "github.com/micmonay/UPnP"
	"github.com/ninnemana/gohbridge/hue/bridge"
	"github.com/pkg/errors"
	"google.golang.org/grpc/grpclog"
)

// Service implements bridge.Interactor around the Hue
// Rest API.
type Service struct {
	clients map[string]string
	Log     grpclog.LoggerV2
	Trace   *trace.Client
}

// Discover retrieves any available Hue Bridge(s) available to the server.
func (s Service) Discover(params *bridge.DiscoverParams, serv bridge.Service_DiscoverServer) error {
	child := s.Trace.NewSpan("hue.discover")
	defer child.Finish()

	client := http.Client{
		Timeout: time.Second * 5,
	}

	var discoverEndpoint string
	switch strings.ToLower(params.GetMethod()) {
	case "upnp":
		up := upnp.NewUPNP(upnp.SERVICE_GATEWAY_IPV4_V2)
		Interface, err := upnp.GetInterfaceByName("en0")
		if err != nil {
			return err
		}

		// get all devices compatible for the service name (timeout 1 second)
		devices := up.GetAllCompatibleDevice(Interface, 1)
		if len(devices) == 0 {
			return errors.Errorf("no devices found on network")
		}

		for _, d := range devices {
			for _, serv := range d.GetAllService() {
				fmt.Println(serv.ControlURL, serv.EventSubURL, serv.SCPDURL)
			}
		}

	case "remote":
		discoverEndpoint = "https://www.meethue.com/api/nupnp"
	default:
		return errors.Errorf("connection method '%s' was not valid", params.GetMethod())
	}

	req, err := http.NewRequest(http.MethodGet, discoverEndpoint, nil)
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

// GetBridgeState retrieves the current state of the configured Hue Bridge.
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

// GetConfig retrieves the full configuration of the requested Hue Bridge.
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
