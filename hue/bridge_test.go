package hue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/google/uuid"
)

var (
	testUser = "6P1KtzwOPY0aiDHOVU4jx7Mn4oPNTqhi6v81hSbG"
)

const (
	// BadJSON is a failure scenario for sending bad data
	BadJSON = "bad_json"
	// NotFound is a failure scenario for sending bad data
	NotFound = "not_found"
	// BadUsername is a failure scenario for sending bad data
	BadUsername = "bad_userna"
)

type DiscoverHandler struct {
	sync.Mutex

	Address string
	Fail    string
}

func (h *DiscoverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var obj interface{}
	switch h.Fail {
	case NotFound:
		w.WriteHeader(404)

		return
	case BadJSON:
		fmt.Println(BadJSON)
		obj = []BridgeState{}
	default:
		obj = []Bridge{
			Bridge{
				ID:   uuid.New().String(),
				User: "discover",
				BridgeNetwork: BridgeNetwork{
					InternalIP: h.Address,
				},
			},
		}
	}
	js, err := json.Marshal(obj)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(js)

}

func TestDiscover(t *testing.T) {

	portalURL = "asl;dfj"
	_, err := Discover()
	if err == nil {
		t.Fatal("error should not be nil")
	}

	handler := &DiscoverHandler{
		Fail: NotFound,
	}
	server := httptest.NewServer(handler)
	defer server.Close()
	portalURL = server.URL

	_, err = Discover()
	if err == nil {
		t.Fatal("error should not be nil")
		t.FailNow()
	}

	portalURL = "http://www.ninneman.org"
	_, err = Discover()
	if err == nil {
		t.Fatal("error should not be nil")
	}

	handler.Fail = ""
	server = httptest.NewServer(handler)
	defer server.Close()
	portalURL = server.URL
	_, err = Discover()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetState(t *testing.T) {

	handler := &DiscoverHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()
	portalURL = server.URL

	ll, err := Discover()
	if err != nil {
		t.Fatal(err)
	}

	if len(ll) == 0 {
		t.Log("no bridges found")
		return
	}

	ip := ll[0].InternalIP
	ll[0].InternalIP = "ftp:[aklsjdf]"
	_, err = ll[0].GetState()
	if err == nil {
		t.Fatal("should have errored on bad username")
	}
	ll[0].InternalIP = ip

	ll[0].User = "random_user"

	_, err = ll[0].GetState()
	if err == nil {
		t.Fatal(err)
	}
}

func TestDo(t *testing.T) {

	br := BridgeRequest{
		Bridge: Bridge{
			BridgeNetwork: BridgeNetwork{
				InternalIP: "localhost:9999",
			},
		},
		Request: &http.Request{},
	}

	_, err := br.Do()
	if err == nil {
		t.Fatal("should error on empty request")
	}

	_, err = br.Bridge.NewRequest("GET", "/", nil, false).Do()
	if err == nil {
		t.Fatal("should error on empty request")
	}
}
