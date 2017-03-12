package hue

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/pkg/errors"
)

// Light the attributes and state of a given light.
type Light struct {
	State             LightState `json:"state"`
	Type              string     `json:"type"`
	Name              string     `json:"name"`
	ModelID           string     `json:"modelid"`
	ID                string     `json:"uniqueid"`
	ManufacturerName  string     `json:"manufactuername"`
	LuminaireIniqueID string     `json:"luminaireuniqueid"`
	ProductID         string     `json:"productid"`
	SWVersion         string     `json:"swversion"`
	SWConfig          string     `json:"swconfig"`

	// This parameter is reserved for future functionality. As from 1.11 point symbols are no longer returned.
	PointSymbol interface{} `json:"pointsymbol"`
}

// LightState the state object contains the following fields.
type LightState struct {
	On               bool      `json:"on"`
	Brightness       uint8     `json:"bri"`
	Hue              uint16    `json:"hue"`
	Saturation       uint8     `json:"sat"`
	ColorCoordinates []float32 `json:"xy"`
	ColorTemperature uint16    `json:"ct"`
	Alert            string    `json:"alert"`
	Effect           string    `json:"effect"`
	ColorMode        string    `json:"colormode"`
	Reachable        bool      `json:"reachable"`
}

// NewLight defines the initial declaration of a light once it receives power
// for the first time.
type NewLight struct {
	Name string `json:"name"`
}

// NewLightList is the list of new, un-configured lights.
type NewLightList map[string]NewLight

// NewLightScan defines the result of a network scan for new lights.
type NewLightScan struct {
	NewLightList
	LastScan string `json:"lastscan"`
}

// SearchRequest is the body of the request being sent on a InitLightSearch
// call.
type SearchRequest struct {
	DeviceIDs []string `json:"deviceid"`
}

// SearchResponse is the body of the response from a InitLightSearch call.
type SearchResponse struct {
	Success interface{} `json:"success"`
}

// GetLights retrieves all lights in the Bridge's network.
func GetLights(b Bridge) (map[string]Light, error) {
	resp, err := b.NewRequest("GET", "/lights", nil).Do()
	if err != nil {
		return nil, err
	}

	switch resp.(type) {
	case map[string]Light:
		return resp.(map[string]Light), nil
	case []interface{}:
		bridgeErr, ok := readError(resp)
		if !ok {
			goto genError
		}

		return nil, errors.Errorf("failed to get lights: %s", bridgeErr[0].Error.Description)
	genError:
		return nil, errors.New("failed to get lights")
	default:
		return nil, errors.New("failed to get lights")
	}

	// return lights, err
}

// GetNewLights retrieves any new, un-configured lights on the Bridge's network.
func GetNewLights(b Bridge) (*NewLightScan, error) {
	obj, err := b.NewRequest("GET", "/lights/new", nil).Do()
	if err != nil {
		return nil, err
	}

	switch obj.(type) {
	case NewLightScan:
		state := obj.(NewLightScan)
		return &state, nil
	case []interface{}:
		bridgeErr, ok := readError(obj)
		if !ok {
			goto genError
		}

		return nil, errors.Errorf("failed to get new lights: %s", bridgeErr[0].Error.Description)
	genError:
		return nil, errors.New("failed to get new lights")
	default:
		return nil, errors.New("failed to get new lights")
	}
}

// InitLightSearch  starts a search for new lights. As of 1.3 will also
// find switches (e.g. "tap")
//
// The bridge will open the network for 40s. The overall search might take
// longer since the configuration of (multiple) new devices can take longer.
// If many devices are found the command will have to be issued a second time
// after discovery time has elapsed. If the command is received again during
// search the search will continue for at least an additional 40s.
//
// When the search has finished, new lights will be available using the get new
// lights command. In addition, the new lights will now be available by calling
// get all lights or by calling get group attributes on group 0. Group 0 is a
// special group that cannot be deleted and will always contain all lights known
// by the bridge.
func InitLightSearch(b Bridge, ids []string) (*SearchResponse, error) {
	var body []byte
	if len(ids) > 0 {
		req := SearchRequest{
			DeviceIDs: ids,
		}
		body, _ = json.Marshal(req)
	}

	obj, err := b.NewRequest("POST", "/lights", bytes.NewBuffer(body)).Do()
	if err != nil {
		return nil, err
	}

	switch obj.(type) {
	case SearchResponse:
		state := obj.(SearchResponse)
		return &state, nil
	case []interface{}:
		bridgeErr, ok := readError(obj)
		if !ok {
			goto genError
		}

		return nil, errors.Errorf("failed to initialize search: %s", bridgeErr[0].Error.Description)
	genError:
		return nil, errors.New("failed to initialize search")
	default:
		return nil, errors.New("failed to initialize search")
	}
}

// GetLight gets the attributes and state of a given light.
func GetLight(b Bridge, id string) (*Light, error) {
	obj, err := b.NewRequest("GET", fmt.Sprintf("/lights/%s", id), nil).Do()
	if err != nil {
		return nil, err
	}

	switch obj.(type) {
	case Light:
		state := obj.(Light)
		return &state, nil
	case []interface{}:
		bridgeErr, ok := readError(obj)
		if !ok {
			goto genError
		}

		return nil, errors.Errorf("failed to get light: %s", bridgeErr[0].Error.Description)
	genError:
		return nil, errors.New("failed to get light")
	default:
		return nil, errors.New("failed to get light")
	}
}

// RenameLight gets the attributes and state of a given light.
func RenameLight(b Bridge, id, name string) error {
	if id == "" {
		return fmt.Errorf("id cannot be blank")
	}
	if name == "" {
		return fmt.Errorf("name cannot be blank")
	}

	body := []byte(fmt.Sprintf("{\"name\": \"%s\"}", name))

	data, err := b.NewRequest("PUT", fmt.Sprintf("/lights/%s", id), bytes.NewBuffer(body)).Do()
	if err != nil {
		return err
	}

	fmt.Println(data)
	return nil
	// var l []LightSearchResponse
	// return json.Unmarshal(data, &l)
}
