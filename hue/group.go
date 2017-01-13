package hue

import "time"

// Group defines a coordinate group a lights as defined by the user.
type Group struct {
	Name         string      `json:"name"`
	LightIndexes []string    `json:"lights"`
	Type         string      `json:"type"`
	ModelID      string      `json:"modelid"`
	ID           string      `json:"uniqueid"`
	State        GroupState  `json:"state"`
	Class        string      `json:"class"`
	Action       GroupAction `json:"action"`
	Recycle      bool        `json:"recycle"`
}

// GroupState is the representation of the current properties of a group of
// lights.
type GroupState struct {
	AllOn           bool        `json:"all_on"`
	AnyOn           bool        `json:"any_on"`
	Brightness      uint8       `json:"bri"`
	LastUpdated     *time.Time  `json:"lastupdated"`
	LastSwitched    *time.Time  `json:"lastswitched"`
	TransmitSymbol  interface{} `json:"transmitsymbol"`
	Duration        uint16      `json:"duration"`
	SymbolSelection string      `json:"symbolselection"`
}

// GroupAction for groups we have the “action” resource (similar to the state
// resource for lights but containing the last values sent to the group
// rather than a state).
type GroupAction struct {
	On                        bool      `json:"on"`
	Brightness                uint8     `json:"bri"`
	Hue                       uint16    `json:"hue"`
	Saturation                uint8     `json:"sat"`
	ColorCoordinates          []float32 `json:"xy"`
	ColorTemperature          uint16    `json:"ct"`
	Alert                     string    `json:"alert"`
	Effect                    string    `json:"effect"`
	TransitionTime            uint16    `json:"transitiontime"`
	BrightnessIncrement       uint8     `json:"bri_inc"`
	SaturationIncrement       uint8     `json:"sat_inc"`
	ColorTemperatureIncrement uint16    `json:"ct_inc"`
	ColorCoordinatesIncrement uint16    `json:"xy_inc"`
	Scene                     string    `json:"scene"`
}
