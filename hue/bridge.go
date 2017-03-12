package hue

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

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
	Bridge   Bridge
	Request  *http.Request
	SkipAuth bool
	Error    error
}

// Bridge defines a unique Philips Hue Bridge.
type Bridge struct {
	BridgeNetwork
	ID   string `json:"id" xml:"id"`
	User string `json:"user,omitempty" xml:"user"`
}

// ErrorResponse wraps the returned error.
type ErrorResponse struct {
	Error BridgeError `json:"error"`
}

// BridgeError defines the error object coming back from the Hue Bridge.
type BridgeError struct {
	Address     string  `json:"address"`
	Description string  `json:"description"`
	Type        float64 `json:"type"`
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
	obj, err := b.NewRequest("GET", "", nil, false).Do()
	if err != nil {
		return nil, err
	}

	bs := &BridgeState{}
	errResp := readJSON(bs, obj)
	switch errResp {
	case nil:
		return bs, nil
	default:
		return nil, errors.Errorf("failed to fetch bridge state: %s", errResp.Error.Description)
	}
}

// NewRequest creates a new request that is bound to the assigned bridge.
func (b Bridge) NewRequest(method, uri string, body io.Reader, skip bool) *BridgeRequest {
	br := BridgeRequest{
		Bridge:   b,
		SkipAuth: skip,
	}

	br.Request, _ = http.NewRequest(method, b.toURI(uri, skip), body)

	return &br
}

// Do executes the BridgeRequest and returns the response body or any error
// that has occurred.
func (br BridgeRequest) Do() ([]byte, error) {
	if !br.SkipAuth && br.Bridge.User == "" {
		return nil, errors.Errorf("no user has been provided")
	}

	client := &http.Client{}
	resp, err := client.Do(br.Request)
	if err != nil {
		return nil, err
	}
	if resp.Body == nil {
		return nil, errors.New("failed to parse response")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	switch {
	case resp.StatusCode != 200 && err == nil:
		return nil, errors.Errorf("%s", data)
	case resp.StatusCode != 200:
		return nil, errors.New("failed to make request")
	default:
		return data, nil
	}
}

func (b Bridge) toURI(route string, skip bool) string {
	switch skip {
	case true:
		return fmt.Sprintf("http://%s/api%s", b.InternalIP, route)
	default:
		return fmt.Sprintf("http://%s/api/%s%s", b.InternalIP, b.User, route)
	}

}

func readJSON(ifc interface{}, data []byte) *ErrorResponse {
	err := json.Unmarshal(data, ifc)
	switch err {
	case nil:
		return nil
	default:
		er := []ErrorResponse{}
		err = json.Unmarshal(data, &er)
		if err != nil {
			er = []ErrorResponse{
				ErrorResponse{
					Error: BridgeError{
						Address:     "",
						Type:        0,
						Description: err.Error(),
					},
				},
			}
		}

		return &er[0]
	}
}

func readError(errors interface{}) ([]ErrorResponse, bool) {

	var vals []interface{}
	switch errors.(type) {
	case []interface{}:
		vals = errors.([]interface{})
	case interface{}:
		vals = []interface{}{errors.(interface{})}
	default:
		return nil, false
	}

	berrs := make([]ErrorResponse, 0)
	for _, err := range vals {
		var errMap map[string]interface{}
		switch err.(type) {
		case map[string]interface{}:
			errMap = err.(map[string]interface{})
		default:
			continue
		}

		errDef, ok := errMap["error"]
		if !ok {
			continue
		}

		switch errDef.(type) {
		case map[string]interface{}:
			errMap = errDef.(map[string]interface{})
		default:
			continue
		}

		var be BridgeError

		// get address
		switch errMap["address"].(type) {
		case string:
			be.Address = errMap["address"].(string)
		}

		// get description
		switch errMap["description"].(type) {
		case string:
			be.Description = errMap["description"].(string)
		}

		// get type
		switch errMap["type"].(type) {
		case int:
			be.Type = errMap["type"].(float64)
		case float32:
			be.Type = errMap["type"].(float64)
		case float64:
			be.Type = errMap["type"].(float64)
		}

		berrs = append(berrs, ErrorResponse{Error: be})
	}

	return berrs, (len(berrs) == len(vals))
}
