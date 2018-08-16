package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"cloud.google.com/go/trace"
	"github.com/ninnemana/huego"
	"github.com/pkg/errors"
)

func (c *client) AllLights(ctx context.Context) ([]interface{}, error) {
	span := trace.FromContext(ctx).NewChild("hue.http.lights.all")
	defer span.Finish()

	user, ok := ctx.Value(hue.UserKey{}).(string)
	if !ok {
		return nil, hue.ErrNoUser
	}

	host, ok := ctx.Value(hue.HostKey{}).(string)
	if !ok {
		return nil, hue.ErrNoHost
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	path := fmt.Sprintf("%s/api/%s/lights", host, user)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(data))
	}

	lights := make(map[string]interface{}, 0)
	err = json.NewDecoder(resp.Body).Decode(&lights)
	if err != nil {
		return nil, err
	}

	results := []interface{}{}
	for key, l := range lights {
		id, err := strconv.Atoi(key)
		if err != nil {
			return nil, errors.Errorf("failed to parse light key into identifier '%s'", key)
		}

		lmp, ok := l.(map[string]interface{})
		if !ok {
			continue
		}

		lmp["ID"] = int32(id)
		results = append(results, lmp)
	}

	return results, nil
}

func (c *client) NewLights(ctx context.Context) (interface{}, error) {
	span := trace.FromContext(ctx).NewChild("hue.http.lights.new")
	defer span.Finish()

	user, ok := ctx.Value(hue.UserKey{}).(string)
	if !ok {
		return nil, hue.ErrNoUser
	}

	host, ok := ctx.Value(hue.HostKey{}).(string)
	if !ok {
		return nil, hue.ErrNoHost
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	req, err := http.NewRequest(
		http.MethodGet,
		fmt.Sprintf("%s/api/%s/lights/new", host, user),
		nil,
	)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(string(data))
	}

	var scan map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&scan)
	if err != nil {
		return nil, err
	}

	return scan, nil
}

func (c *client) SearchLights(ctx context.Context, deviceIDs []string) error {
	span := trace.FromContext(ctx).NewChild("hue.http.lights.search")
	defer span.Finish()

	user, ok := ctx.Value(hue.UserKey{}).(string)
	if !ok {
		return hue.ErrNoUser
	}

	host, ok := ctx.Value(hue.HostKey{}).(string)
	if !ok {
		return hue.ErrNoHost
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	var buf io.Reader
	if len(deviceIDs) > 0 {
		js, err := json.Marshal(hue.SearchParams{
			Devices: deviceIDs,
		})
		if err != nil {
			return err
		}

		buf = bytes.NewBuffer(js)
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/%s/lights", host, user),
		buf,
	)
	if err != nil {
		return err
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(data))
	}

	return nil
}

func (c *client) GetLight(ctx context.Context, id int) (interface{}, error) {
	span := trace.FromContext(ctx).NewChild("hue.http.lights.get")
	defer span.Finish()

	user, ok := ctx.Value(hue.UserKey{}).(string)
	if !ok {
		return nil, hue.ErrNoUser
	}

	host, ok := ctx.Value(hue.HostKey{}).(string)
	if !ok {
		return nil, hue.ErrNoHost
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	path := fmt.Sprintf("%s/api/%s/lights/%d", host, user, id)
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(data))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var l map[string]interface{}
	if err := json.Unmarshal(data, &l); err != nil {
		return nil, errors.Errorf("failed to encode '%s' to Light: %v", data, err)
	}

	l["ID"] = int32(id)

	return l, nil
}

func (c *client) RenameLight(string, string) (interface{}, error) {
	return nil, hue.ErrNotImplemented
}

func (c *client) LightState(ctx context.Context, id int, state interface{}) (interface{}, error) {
	span := trace.FromContext(ctx).NewChild("hue.http.lights.state")
	defer span.Finish()

	user, ok := ctx.Value(hue.UserKey{}).(string)
	if !ok {
		return nil, hue.ErrNoUser
	}

	host, ok := ctx.Value(hue.HostKey{}).(string)
	if !ok {
		return nil, hue.ErrNoHost
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	data, err := json.Marshal(state)
	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("%s/api/%s/lights/%d/state", host, user, id)
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(data))
	}

	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var res []interface{}
	if err = json.Unmarshal(data, &res); err != nil {
		return nil, err
	}

	for _, r := range res {
		switch r.(type) {
		case map[string]interface{}:
			result := r.(map[string]interface{})
			switch {
			case result["error"] != nil:
				switch result["error"].(type) {
				case map[string]interface{}:
					e := result["error"].(map[string]interface{})
					switch e["description"].(type) {
					case string:
						return nil, errors.Errorf("failed to set state: %s", e["description"])
					}
				}

				return nil, errors.Errorf("state update failed")
			}
		default:
			return nil, errors.Errorf("failed to read state response")
		}
	}

	return c.GetLight(ctx, id)
}

func (c *client) Toggle(ctx context.Context, id int) (interface{}, error) {
	span := trace.FromContext(ctx).NewChild("hue.http.lights.toggle")
	defer span.Finish()

	user, ok := ctx.Value(hue.UserKey{}).(string)
	if !ok {
		return nil, hue.ErrNoUser
	}

	host, ok := ctx.Value(hue.HostKey{}).(string)
	if !ok {
		return nil, hue.ErrNoHost
	}

	res, err := c.GetLight(ctx, id)
	if err != nil {
		return nil, err
	}

	existing, ok := res.(map[string]interface{})
	if !ok {
		return nil, errors.Errorf("failed to convert '%T' to *light.Light", res)
	}

	alreadyOn := false
	existingMap, ok := existing["state"].(map[string]interface{})
	if !ok {
		return nil, errors.Errorf("existing could not be mapped from '%T' to map[string]interface{}", existing)
	}

	on, ok := existingMap["on"].(bool)
	if !ok {
		return nil, errors.Errorf("existing could not be mapped from '%T' to bool", existingMap["on"])
	}

	if on {
		alreadyOn = true
	}

	client := http.Client{
		Timeout: time.Second * 5,
	}

	body := fmt.Sprintf(`{"on": %t}`, !alreadyOn)

	path := fmt.Sprintf("%s/api/%s/lights/%d/state", host, user, id)
	req, err := http.NewRequest(http.MethodPut, path, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return nil, err
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(data))
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []interface{}
	if err = json.Unmarshal(data, &result); err != nil {
		return nil, err
	}

	for _, r := range result {
		switch r.(type) {
		case map[string]interface{}:
			result := r.(map[string]interface{})
			switch {
			case result["error"] != nil:
				switch result["error"].(type) {
				case map[string]interface{}:
					e := result["error"].(map[string]interface{})
					switch e["description"].(type) {
					case string:
						return nil, errors.Errorf("failed to set state: %s", e["description"])
					}
				}

				return nil, errors.Errorf("state update failed")
			}
		default:
			return nil, errors.Errorf("failed to read state response")
		}
	}

	return c.GetLight(ctx, id)
}

func (c *client) DeleteLight(string) error {
	return hue.ErrNotImplemented
}
