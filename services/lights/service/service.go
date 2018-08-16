package service

import (
	"context"
	"fmt"

	light "github.com/ninnemana/gohbridge/services/lights"
	"github.com/ninnemana/huego"

	"github.com/golang/protobuf/ptypes/empty"
	jsoniter "github.com/ninnemana/json-iterator"
	"github.com/pkg/errors"
	"go.opencensus.io/trace"
)

var (
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

// Service implements lights.Interactor
type Service struct {
	hue hue.Client
}

// New instantiates a new implementation of the light gRPC interface.
func New(cl hue.Client) (*Service, error) {
	if cl == nil {
		return nil, errors.New("required Philips Hue client was no supplied")
	}

	return &Service{
		hue: cl,
	}, nil
}

// All retrieves a streamed list of light records retrieved from the Hue REST API.
func (s *Service) All(params *light.ListParams, server light.Service_AllServer) error {
	_, span := trace.StartSpan(server.Context(), "hue.lights.all")
	defer span.End()

	ctx := context.WithValue(server.Context(), hue.UserKey{}, params.GetUser())
	ctx = context.WithValue(ctx, hue.HostKey{}, params.GetHost())

	results, err := s.hue.AllLights(ctx)
	if err != nil {
		return err
	}

	for _, r := range results {
		l := &light.Light{}
		switch r.(type) {
		case *light.Light:
			l = r.(*light.Light)
		case light.Light:
			t := r.(light.Light)
			l = &t
		case map[string]interface{}:
			data, err := json.Marshal(r)
			if err != nil {
				return err
			}

			if err := json.Unmarshal(data, l); err != nil {
				return err
			}
		default:
			return errors.Errorf("failed to convert '%T' into *light.Light", r)
		}

		if err = server.Send(l); err != nil {
			return err
		}
	}

	return nil
}

// New polls the Hue Bridge for lights that are available to be setup.
func (s *Service) New(ctx context.Context, params *light.NewParams) (*light.Scan, error) {
	_, span := trace.StartSpan(ctx, "hue.lights.new")
	defer span.End()

	ctx = context.WithValue(ctx, hue.UserKey{}, params.GetUser())
	ctx = context.WithValue(ctx, hue.HostKey{}, params.GetHost())

	res, err := s.hue.NewLights(ctx)
	if err != nil {
		return nil, err
	}

	scan, ok := res.(*light.Scan)
	if !ok {
		return nil, errors.Errorf("failed to convert '%T' to *lights.Scan", res)
	}

	return scan, nil
}

// Search initializes a light search event.
func (s *Service) Search(ctx context.Context, params *light.SearchParams) (*empty.Empty, error) {
	_, span := trace.StartSpan(ctx, "hue.lights.search")
	defer span.End()

	ctx = context.WithValue(ctx, hue.UserKey{}, params.GetUser())
	ctx = context.WithValue(ctx, hue.HostKey{}, params.GetHost())

	return nil, s.hue.SearchLights(ctx, params.GetDevices())
}

// Get retrieves information about a specific light.
func (s *Service) Get(ctx context.Context, params *light.GetParams) (*light.Light, error) {
	_, span := trace.StartSpan(ctx, "hue.lights.get")
	defer span.End()

	ctx = context.WithValue(ctx, hue.UserKey{}, params.GetUser())
	ctx = context.WithValue(ctx, hue.HostKey{}, params.GetHost())

	res, err := s.hue.GetLight(ctx, int(params.GetID()))
	if err != nil {
		return nil, err
	}

	var l *light.Light
	switch res.(type) {
	case *light.Light:
		var ok bool
		l, ok = res.(*light.Light)
		if !ok {
			return nil, errors.Errorf("failed to convert '%T' to *light.Light", res)
		}
	case map[string]interface{}:
		mp, ok := res.(map[string]interface{})
		if !ok {
			return nil, errors.Errorf("failed to convert '%T' to map[string]interface{}", res)
		}

		data, err := json.Marshal(mp)
		if err != nil {
			return nil, errors.Wrap(err, "failed to marshal into JSON data")
		}

		if err := json.Unmarshal(data, &l); err != nil {
			return nil, errors.Wrap(err, "failed to unmarshal JSON data into *light.Light")
		}

		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeFailedPrecondition,
			Message: "sample error",
		})
		fmt.Println("set status")
	default:
		return nil, errors.Errorf("failed to convert '%T' to *light.Light", res)
	}

	return l, nil
}

// SetState updates the current state of the defined light.
func (s *Service) SetState(ctx context.Context, params *light.SetStateParams) (*light.Light, error) {
	_, span := trace.StartSpan(ctx, "hue.lights.state.set")
	defer span.End()

	ctx = context.WithValue(ctx, hue.UserKey{}, params.GetUser())
	ctx = context.WithValue(ctx, hue.HostKey{}, params.GetHost())

	res, err := s.hue.LightState(ctx, int(params.GetID()), params.GetUpdate())
	if err != nil {
		return nil, err
	}

	l, ok := res.(*light.Light)
	if !ok {
		return nil, errors.Errorf("failed to convert '%T' to *light.Light", res)
	}

	return l, nil
}

// Toggle changes the light from on -> off or off -> on depending on the current state of the light.
func (s *Service) Toggle(ctx context.Context, params *light.ToggleParams) (*light.Light, error) {
	_, span := trace.StartSpan(ctx, "hue.lights.toggle")
	defer span.End()

	ctx = context.WithValue(ctx, hue.UserKey{}, params.GetUser())
	ctx = context.WithValue(ctx, hue.HostKey{}, params.GetHost())

	res, err := s.hue.Toggle(ctx, int(params.GetID()))
	if err != nil {
		return nil, err
	}

	l, ok := res.(*light.Light)
	if !ok {
		return nil, errors.Errorf("failed to convert '%T' to *light.Light", res)
	}

	return l, nil
}
