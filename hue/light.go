package hue

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
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

// LightSearchRequest is the body of the request being sent on a InitLightSearch
// call.
type LightSearchRequest struct {
	DeviceIDs []string `json:"deviceid"`
}

// LightSearchResponse is the body of the response from a InitLightSearch call.
type LightSearchResponse struct {
	Success map[string]string `json:"success"`
}

// GetLights retrieves all lights in the Bridge's network.
func GetLights(b Bridge) (map[string]Light, error) {
	data, err := b.NewRequest("GET", fmt.Sprintf("%s/lights", os.Getenv("HUE_USER")), nil).Do()
	if err != nil {
		return nil, err
	}

	var lights map[string]Light
	err = json.Unmarshal(data, &lights)

	return lights, err
}

// GetNewLights retrieves any new, un-configured lights on the Bridge's network.
func GetNewLights(b Bridge) (*NewLightScan, error) {
	data, err := b.NewRequest("GET", fmt.Sprintf("%s/lights/new", os.Getenv("HUE_USER")), nil).Do()
	if err != nil {
		return nil, err
	}

	var lights NewLightScan
	err = json.Unmarshal(data, &lights)

	return &lights, err
}

// InitLightSearch opens the bridge for 40s. The overall search might take
// longer since the configuration of (multiple) new devices can take longer.
// If many devices are found the command will have to be issued a second time
// after discovery time has elapsed. If the command is received again during
// search the search will continue for at least an additional 40s.
//
// When the search has finished, new lights will be available using
// `GetNewLights`. In addition, the new lights will now be available by
// calling `GetLights` or by calling get group attributes on group 0.
// Group 0 is a special group that cannot be deleted and will always contain
// all lights known by the bridge.
func InitLightSearch(b Bridge, ids ...string) error {
	var body []byte
	if len(ids) > 0 {
		req := LightSearchRequest{
			DeviceIDs: ids,
		}
		body, _ = json.Marshal(req)
	}
	data, err := b.NewRequest("POST", fmt.Sprintf("%s/lights", os.Getenv("HUE_USER")), bytes.NewBuffer(body)).Do()
	if err != nil {
		return err
	}

	var resp LightSearchResponse
	err = json.Unmarshal(data, &resp)

	return err
}

// GetLight gets the attributes and state of a given light.
func GetLight(b Bridge, id string) (*Light, error) {
	data, err := b.NewRequest("GET", fmt.Sprintf("%s/lights/%s", os.Getenv("HUE_USER"), id), nil).Do()
	if err != nil {
		return nil, err
	}

	var l *Light
	err = json.Unmarshal(data, &l)

	return l, err
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

	data, err := b.NewRequest("PUT", fmt.Sprintf("%s/lights/%s", os.Getenv("HUE_USER"), id), bytes.NewBuffer(body)).Do()
	if err != nil {
		return err
	}

	var l LightSearchResponse
	return json.Unmarshal(data, &l)
}
