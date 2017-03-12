package api

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ninnemana/gohbridge/hue"
	"github.com/pkg/errors"
)

const (
	contextKeyToken = contextKey("api-service")
)

type contextKey string

type Service struct {
	Address         string
	User            string
	Path            string
	RequestBody     interface{}
	RequestHeaders  map[string][]string
	ResponseBody    interface{}
	ResponseHeaders map[string][]string
	Bridge          *hue.Bridge
}

// WithService is our middleware wrapper that authenticates an API request
// and scopes the service to the request context.
func WithService(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var s Service

		bridges, err := hue.Discover()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}

		if len(bridges) == 0 {
			http.Error(w, "no bridges found", http.StatusBadGateway)
			return
		}

		s.Bridge = &bridges[0]
		err = s.auth(r.Header)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		s.Bridge.User = s.User

		ctx := context.WithValue(r.Context(), contextKeyToken, s)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (s *Service) auth(h http.Header) error {
	val := h.Get("Authorization")
	if val == "" {
		return errors.Errorf("missing authorization header")
	}

	s.User = val
	s.Bridge.User = val

	return nil
}

func (c contextKey) String() string {
	return "api-context-key " + string(c)
}

func parseBody(s Service, body io.Reader) (interface{}, error) {
	data, err := ioutil.ReadAll(body)

	contentType := s.RequestHeaders["Content-Type"]
	switch strings.ToLower(contentType[0]) {
	case "application/json":
		fmt.Println(data, err)
	}

	return nil, nil
}
