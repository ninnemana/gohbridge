package bridge

import (
	"context"

	pb "github.com/ninnemana/gohbridge/hue/bridge/service"
)

// BridgeInteractor describes the functionality of the Bridge service.
type BridgeInteractor interface {
	Discover(*pb.DiscoverParams, pb.Hue_DiscoverServer) error
	GetBridgeState(context.Context, *pb.Bridge) (*pb.BridgeState, error)
}
