package hue

import "context"

type AllBridgeParams struct {
	Method string
}

type SearchParams struct {
	Devices []string `json:"deviceid"`
}

type UserKey struct{}
type HostKey struct{}

type Client interface {
	AllLights(context.Context) ([]interface{}, error)
	NewLights(context.Context) (interface{}, error)
	SearchLights(context.Context, []string) error
	GetLight(context.Context, int) (interface{}, error)
	RenameLight(string, string) (interface{}, error)
	LightState(context.Context, int, interface{}) (interface{}, error)
	Toggle(context.Context, int) (interface{}, error)
	DeleteLight(string) error

	AllGroups() ([]interface{}, error)
	CreateGroup(interface{}) (interface{}, error)
	GetGroup(string) (interface{}, error)
	SaveGroup(string, interface{}) (interface{}, error)
	SetGroupState(string, interface{}) (interface{}, error)
	DeleteGroup(string) error

	AllSchedules() ([]interface{}, error)
	CreateSchedule(interface{}) (interface{}, error)
	GetSchedule(string) (interface{}, error)
	SetSchedule(string, interface{}) (interface{}, error)
	DeleteSchedule(string) error

	AllScenes() ([]interface{}, error)
	GetScene(string) (interface{}, error)
	CreateScene(interface{}) (interface{}, error)
	SetScene(string, interface{}) (interface{}, error)
	DeleteScene(string) error

	AllSensors() ([]interface{}, error)
	CreateSensor(interface{}) (interface{}, error)
	SearchSensors() error
	NewSensors() ([]interface{}, error)
	GetSensor(string) (interface{}, error)
	SetSensor(string, interface{}) (interface{}, error)
	RenameSensor(string, string) (interface{}, error)
	DeleteSensor(string) error

	AllRules() ([]interface{}, error)
	GetRule(string) (interface{}, error)
	CreateRule(interface{}) (interface{}, error)
	UpdateRule(string, interface{}) (interface{}, error)
	DeleteRule(string) error

	AllBridges(context.Context, interface{}) ([]interface{}, error)
	CreateUser(interface{}) (interface{}, error)
	GetConfig(context.Context) (interface{}, error)
	ModifyConfig(interface{}) (interface{}, error)
	Unwhitelist(string) error
	GetFullState(context.Context) (interface{}, error)
}

type Bridge struct {
	ID                string `json:"id"`
	InternalIPAddress string `json:"internalipaddress"`
	MacAddress        string `json:"macaddress"`
}
