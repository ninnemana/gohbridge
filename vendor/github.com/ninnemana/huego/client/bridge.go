package client

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/ninnemana/huego"

	jsoniter "github.com/ninnemana/json-iterator"
	"github.com/pkg/errors"
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func (c *client) AllBridges(ctx context.Context, q interface{}) ([]interface{}, error) {
	span := trace.NewSpan(
		"hue.http.bridges.all",
		trace.FromContext(ctx),
		trace.StartOptions{},
	)
	defer span.End()

	params, ok := q.(*hue.AllBridgeParams)
	if !ok {
		return nil, errors.Errorf("provided params were expected to be *AllBridgeParams, received '%T'", params)
	}

	// Automatically add a Stackdriver trace header to outgoing requests:
	client := &http.Client{
		Transport: &ochttp.Transport{},
	}

	span.AddAttributes(trace.StringAttribute("method", params.Method))
	var discoverEndpoint string
	switch strings.ToLower(params.Method) {
	// case "upnp":
	// up := upnp.NewUPNP(upnp.SERVICE_GATEWAY_IPV4_V2)
	// Interface, err := upnp.GetInterfaceByName("en0")
	// if err != nil {
	// 	return nil, err
	// }

	// // get all devices compatible for the service name (timeout 1 second)
	// devices := up.GetAllCompatibleDevice(Interface, 1)
	// if len(devices) == 0 {
	// 	return nil, errors.Errorf("no devices found on network")
	// }

	// for _, d := range devices {
	// 	for _, serv := range d.GetAllService() {
	// 		fmt.Println(serv.ControlURL, serv.EventSubURL, serv.SCPDURL)
	// 	}
	// }

	case "remote":
		discoverEndpoint = "https://www.meethue.com/api/nupnp"
	default:
		return nil, errors.Errorf("connection method '%s' was not valid", params.Method)
	}

	req, err := http.NewRequest(http.MethodGet, discoverEndpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	span.AddAttributes(trace.Int64Attribute("statusCode", int64(resp.StatusCode)))
	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(data))
	}

	bridges := []interface{}{}
	err = json.NewDecoder(resp.Body).Decode(&bridges)
	if err != nil {
		return nil, err
	}

	span.AddAttributes(
		trace.Int64Attribute("count", int64(len(bridges))),
	)

	return bridges, nil
}

func (c *client) CreateUser(interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) GetConfig(ctx context.Context) (interface{}, error) {
	span := trace.NewSpan(
		"hue.http.bridges.config",
		trace.FromContext(ctx),
		trace.StartOptions{},
	)
	defer span.End()

	user, ok := ctx.Value(hue.UserKey{}).(string)
	if !ok {
		return nil, hue.ErrNoUser
	}

	host, ok := ctx.Value(hue.HostKey{}).(string)
	if !ok {
		return nil, hue.ErrNoHost
	}

	path := fmt.Sprintf("%s/api/%s/config", host, user)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: time.Second * 5,
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

	var conf interface{}
	err = json.NewDecoder(resp.Body).Decode(&conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}

func (c *client) ModifyConfig(interface{}) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) Unwhitelist(string) error {
	return hue.ErrNotImplemented
}

func (c *client) GetFullState(ctx context.Context) (interface{}, error) {
	span := trace.NewSpan(
		"hue.http.bridges.state",
		trace.FromContext(ctx),
		trace.StartOptions{},
	)
	defer span.End()

	user, ok := ctx.Value(hue.UserKey{}).(string)
	if !ok {
		return nil, hue.ErrNoUser
	}

	host, ok := ctx.Value(hue.HostKey{}).(string)
	if !ok {
		return nil, hue.ErrNoHost
	}

	path := fmt.Sprintf("%s/api/%s", host, user)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	client := http.Client{
		Timeout: time.Second * 5,
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

	var result interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
