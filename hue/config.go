package hue

import (
	"bytes"
	"encoding/json"

	"github.com/pkg/errors"
)

// Config defines the configuration of the Hub network.
type Config struct {
	Name             string               `json:"name"`             // Name of the bridge. This is also its uPnP name, so will reflect the actual uPnP name after any conflicts have been resolved.
	SWUpdate         SWUpdate             `json:"swupdate"`         // Contains information related to software updates.
	Whitelist        map[string]Whitelist `json:"whitelist"`        // A list of whitelisted user IDs.
	APIVersion       string               `json:"apiversion"`       // The version of the hue API in the format <major>.<minor>.<patch>, for example 1.2.1
	SWVersion        string               `json:"swversion"`        // Software version of the bridge.
	ProxyAddress     string               `json:"proxyaddress"`     // IP Address of the proxy server being used. A value of “none” indicates no proxy.
	ProxyPort        uint16               `json:"proxyport"`        // Port of the proxy being used by the bridge. If set to 0 then a proxy is not being used.
	LinkButton       bool                 `json:"linkbutton"`       // Indicates whether the link button has been pressed within the last 30 seconds.
	IPAddress        string               `json:"ipaddress"`        // IP address of the bridge.
	Mac              string               `json:"mac"`              // MAC address of the bridge.
	Netmask          string               `json:"netmask"`          // Network mask of the bridge.
	Gateway          string               `json:"gateway"`          // Gateway IP address of the bridge.
	DHCP             bool                 `json:"dhcp"`             // Whether the IP address of the bridge is obtained with DHCP.
	PortalServices   bool                 `json:"portalservices"`   // This indicates whether the bridge is registered to synchronize data with a portal account.
	UTC              string               `json:"UTC"`              // Current time stored on the bridge.
	LocalTime        string               `json:"localtime"`        // The local time of the bridge. "none" if not available.
	Timezone         string               `json:"timezone"`         // Timezone of the bridge as OlsenIDs, like "Europe/Amsterdam" or "none" when not configured.
	ZigbeeChannel    uint16               `json:"zigbeechannel"`    // The current wireless frequency channel used by the bridge. It can take values of 11, 15, 20,25 or 0 if undefined (factory new).
	ModelID          string               `json:"modelid"`          // This parameter uniquely identifies the hardware model of the bridge (BSB001, BSB002).
	BridgeID         string               `json:"bridgeid"`         // The unique bridge id. This is currently generated from the bridge Ethernet mac address.
	FactoryNew       bool                 `json:"factorynew"`       // Indicates if bridge settings are factory new.
	ReplacesBridgeID string               `json:"replacesbridgeid"` // If a bridge backup file has been restored on this bridge from a bridge with a different bridgeid, it will indicate that bridge id, otherwise it will be null.
	DatastoreVersion string               `json:"datastoreversion"` // The version of the datastore.
}

// Whitelist is a set of users that are allowed to access this network.
type Whitelist struct {
	LastUseDate string `json:"last use date"`
	CreateDate  string `json:"created date"`
	Name        string `json:"name"`
}

// SWUpdate ...
type SWUpdate struct {
	UpdateState uint8  `json:"updatestate"`
	URL         string `json:"url"`
	Text        string `json:"text"`
	Notify      bool   `json:"notify"`
}

// NewUser
type NewUser struct {
	Username string `json:"username"`
}

type NewUserResponse []struct {
	Success NewUser `json:"success"`
}

type NewUserRequest struct {
	DeviceType string `json:"devicetype"`
}

// CreateUser creates a new user. The link button on the bridge must be
// pressed and this command executed within 30 seconds.
//
// Once a new user has been created, the user key is added to a ‘whitelist’,
// allowing access to API commands that require a whitelisted user. At present,
// all other API commands require a whitelisted user.
//
// All published apps are asked to use the name of their app as the devicetype.
func CreateUser(b Bridge) (*NewUser, error) {

	dr := NewUserRequest{
		DeviceType: b.ID,
	}

	data, err := json.Marshal(dr)
	if err != nil {
		return nil, err
	}

	obj, err := b.NewRequest("POST", "", bytes.NewBuffer(data)).Do()
	if err != nil {
		return nil, err
	}

	switch obj.(type) {
	case NewUserResponse:
		state := obj.(NewUserResponse)
		if len(state) == 0 {
			return nil, errors.New("user was not created")
		}
		return &state[0].Success, nil
	case []interface{}:
		bridgeErr, ok := readError(obj)
		if !ok {
			goto genError
		}

		return nil, errors.Errorf("failed to create user: %s", bridgeErr[0].Error.Description)
	genError:
		return nil, errors.New("failed to create user")
	default:
		return nil, errors.New("failed to create user")
	}
}
