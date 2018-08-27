package bridge

import (
	"context"
)

// Interactor describes the functionality of the Bridge service.
type Interactor interface {
	Discover(*DiscoverParams, Service_DiscoverServer) error
	GetBridgeState(context.Context, *ConfigParams) (*BridgeState, error)
	GetConfig(context.Context, *ConfigParams) (*BridgeConfig, error)
}
