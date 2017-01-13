package hue

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var (
	portalURL     = "https://www.meethue.com/api/nupnp"
	requestMethod = "GET"
)

// BridgeNetwork describes the network properties of a given
// Philips Hue bridge.
type BridgeNetwork struct {
	InternalIP string `json:"internalipaddress" xml:"internalipaddress"`
}

// BridgeRequest ...
type BridgeRequest struct {
	Bridge  Bridge
	Request *http.Request
}

// Bridge defines a unique Philips Hue Bridge.
type Bridge struct {
	BridgeNetwork
	ID string `json:"id" xml:"id"`
}

// BridgeState ...
type BridgeState struct {
	Lights        map[string]Light        `json:"lights"`
	Groups        map[string]Group        `json:"groups"`
	Config        Config                  `json:"config"`
	Schedules     map[string]Schedule     `json:"schedules"`
	Scenes        map[string]Scene        `json:"scenes"`
	Rules         map[string]Rule         `json:"rules"`
	Sensors       map[string]Sensor       `json:"sensors"`
	ResourceLinks map[string]ResourceLink `json:"resourcelinks"`
}

// Discover finds all Hue Bridge on the current network
// by using the meethue.com lookup service.
func Discover() ([]Bridge, error) {

	req, err := http.NewRequest(requestMethod, portalURL, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", string(data))
	}

	var bridges []Bridge
	err = json.Unmarshal(data, &bridges)
	if err != nil {
		return nil, err
	}

	return bridges, nil
}

// GetState ...
func (b Bridge) GetState() (*BridgeState, error) {
	data, err := b.NewRequest("GET", os.Getenv("HUE_USER"), nil).Do()
	if err != nil {
		return nil, err
	}

	var bs *BridgeState
	err = json.Unmarshal(data, &bs)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

// NewRequest ...
func (b Bridge) NewRequest(method, uri string, body io.Reader) *BridgeRequest {
	br := BridgeRequest{
		Bridge: b,
	}
	req, err := http.NewRequest("GET", b.toURI(uri), nil)
	if err != nil {
		return &br
	}

	br.Request = req

	return &br
}

// Do ...
func (br BridgeRequest) Do() ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(br.Request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data []byte
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Hue Request failed: %s", data)
	}

	return data, err
}

func (b Bridge) toURI(route string) string {
	return fmt.Sprintf("http://%s/api/%s", b.InternalIP, route)
}
