package service

import (
	"context"

	"cloud.google.com/go/trace"
	"github.com/ninnemana/gohbridge/hue"
	"github.com/ninnemana/gohbridge/services/bridge"
	jsoniter "github.com/ninnemana/json-iterator"
	"github.com/pkg/errors"
	"google.golang.org/grpc/grpclog"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary

	noHueClient = errors.New("required Philips Hue client was no supplied")
	noLogger    = errors.New("required logger was not provided")
	noTrace     = errors.New("required trace client was not provided")
)

// Service implements bridge.Interactor around the Hue
// Rest API.
type Service struct {
	clients map[string]string
	Log     grpclog.LoggerV2
	Trace   *trace.Client
	hue     hue.Client
}

// New instantiates a new implementation of the bridge gRPC interface.New
func New(cl hue.Client, l grpclog.LoggerV2, tr *trace.Client) (*Service, error) {
	if cl == nil {
		return nil, noHueClient
	}

	if l == nil {
		return nil, noLogger
	}

	if tr == nil {
		return nil, noTrace
	}

	return &Service{
		hue:   cl,
		Log:   l,
		Trace: tr,
	}, nil
}

// Discover retrieves any available Hue Bridge(s) available to the server.
func (s Service) Discover(params *bridge.DiscoverParams, serv bridge.Service_DiscoverServer) error {
	child := s.Trace.NewSpan("hue.discover")
	defer child.Finish()

	res, err := s.hue.AllBridges(serv.Context(), &hue.AllBridgeParams{
		Method: "remote",
	})
	if err != nil {
		return err
	}

	s.Log.Info("remote", res)

	var bridges []*bridge.Bridge
	for _, r := range res {
		switch r.(type) {
		case bridge.Bridge:
			br := r.(bridge.Bridge)
			bridges = append(bridges, &br)
		case *bridge.Bridge:
			bridges = append(bridges, r.(*bridge.Bridge))
		case map[string]interface{}:
			data, err := json.Marshal(r)
			if err != nil {
				return err
			}

			var br bridge.Bridge
			if err := json.Unmarshal(data, &br); err != nil {
				return err
			}

			bridges = append(bridges, &br)
		default:
			return errors.Errorf("failed to parse '%T'", r)
		}
	}

	for _, br := range bridges {
		err = serv.Send(br)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetBridgeState retrieves the current state of the configured Hue Bridge.
func (s Service) GetBridgeState(ctx context.Context, params *bridge.ConfigParams) (*bridge.BridgeState, error) {
	span := trace.FromContext(ctx)
	child := span.NewChild("hue.bridge_state")
	defer child.Finish()

	ctx = context.WithValue(ctx, hue.HostKey{}, params.GetHost())
	ctx = context.WithValue(ctx, hue.UserKey{}, params.GetUser())

	state, err := s.hue.GetFullState(ctx)
	if err != nil {
		return nil, err
	}

	switch state.(type) {
	case bridge.BridgeState:
		bs := state.(bridge.BridgeState)
		return &bs, nil
	case *bridge.BridgeState:
		return state.(*bridge.BridgeState), nil
	case map[string]interface{}:
		data, err := json.Marshal(state)
		if err != nil {
			return nil, err
		}

		var bs bridge.BridgeState
		if err := json.Unmarshal(data, &bs); err != nil {
			return nil, err
		}

		return &bs, nil
	default:
		return nil, errors.Errorf("failed to to parse '%T'", state)
	}
}

// GetConfig retrieves the full configuration of the requested Hue Bridge.
func (s Service) GetConfig(ctx context.Context, params *bridge.ConfigParams) (*bridge.BridgeConfig, error) {
	span := trace.FromContext(ctx)
	child := span.NewChild("hue.full_config")
	defer child.Finish()

	ctx = context.WithValue(ctx, hue.HostKey{}, params.GetHost())
	ctx = context.WithValue(ctx, hue.UserKey{}, params.GetUser())

	cfg, err := s.hue.GetFullState(ctx)
	if err != nil {
		return nil, err
	}

	switch cfg.(type) {
	case bridge.BridgeConfig:
		bc := cfg.(bridge.BridgeConfig)
		return &bc, nil
	case *bridge.BridgeConfig:
		return cfg.(*bridge.BridgeConfig), nil
	case map[string]interface{}:
		data, err := json.Marshal(cfg)
		if err != nil {
			return nil, err
		}

		var bc bridge.BridgeConfig
		if err := json.Unmarshal(data, &bc); err != nil {
			return nil, err
		}

		return &bc, nil
	default:
		return nil, errors.Errorf("failed to to parse '%T'", cfg)
	}
}
