package hue

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"cloud.google.com/go/trace"
	"github.com/pkg/errors"
	context "golang.org/x/net/context"
)

const (
	bridgeState = "/api/<hue-user-id>"
)

type Service struct {
	clients map[string]string
}

func (s Service) Discover(params *DiscoverParams, serv Hue_DiscoverServer) error {
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

	bridges := []Bridge{}
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

func (s Service) GetBridgeState(ctx context.Context, params *Bridge) (*BridgeState, error) {
	span := trace.FromContext(ctx)
	child := span.NewChild("hue.bridge_state")
	defer child.Finish()

	client := http.Client{
		Timeout: time.Second * 5,
	}

	path := fmt.Sprintf("http://192.168.86.133%s", bridgeState)
	fmt.Println(path)
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

	bs := BridgeState{}
	err = json.NewDecoder(resp.Body).Decode(&bs)
	if err != nil {
		return nil, err
	}

	return &bs, nil
}
