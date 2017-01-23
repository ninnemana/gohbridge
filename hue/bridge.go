package hue

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/pkg/errors"
)

var (
	portalURL = "https://www.meethue.com/api/nupnp"
)

// BridgeNetwork describes the network properties of a given
// Philips Hue bridge.
type BridgeNetwork struct {
	InternalIP string `json:"internalipaddress" xml:"internalipaddress"`
}

// BridgeRequest describes a request to be executed against a
// provided Bridge.
type BridgeRequest struct {
	Bridge  Bridge
	Request *http.Request
	Error   error
}

// Bridge defines a unique Philips Hue Bridge.
type Bridge struct {
	BridgeNetwork
	ID string `json:"id" xml:"id"`
}

// BridgeState represents the current logical state of the Hue Bridge.
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

	resp, err := http.Get(portalURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s", "discovery failed")
	}

	var bridges []Bridge
	err = json.NewDecoder(resp.Body).Decode(&bridges)
	if err != nil {
		return nil, err
	}

	return bridges, nil
}

// GetState delivers the returned specific of how the Bridge is feeling, or an
// error if it fails to query the REST API, or parsing the JSON.
func (b Bridge) GetState() (*BridgeState, error) {
	data, err := b.NewRequest("GET", os.Getenv("HUE_USER"), nil).Do()
	if err != nil {
		return nil, err
	}

	var bs *BridgeState
	err = json.Unmarshal(data, &bs)
	return bs, err
}

// NewRequest creates a new request that is bound to the assigned bridge.
func (b Bridge) NewRequest(method, uri string, body io.Reader) *BridgeRequest {
	br := BridgeRequest{
		Bridge: b,
	}

	br.Request, _ = http.NewRequest(method, b.toURI(uri), body)

	return &br
}

// Do executes the BridgeRequest and returns the response body or any error
// that has occurred.
func (br BridgeRequest) Do() ([]byte, error) {
	client := &http.Client{}
	resp, err := client.Do(br.Request)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		data, _ := ioutil.ReadAll(resp.Body)
		return nil, errors.Errorf("Hue Request failed: %s", data)
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (b Bridge) toURI(route string) string {
	return fmt.Sprintf("http://%s/api/%s", b.InternalIP, route)
}
