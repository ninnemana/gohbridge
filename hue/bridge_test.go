package hue

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

type DiscoverHandler struct {
	sync.Mutex
}

func (h *DiscoverHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func TestDiscover(t *testing.T) {

	urlBackup := portalURL
	portalURL = "asl;dfj"
	_, err := Discover()
	if err == nil {
		t.Fatal("error should not be nil")
	}
	portalURL = urlBackup

	handler := &DiscoverHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	urlBackup = portalURL
	portalURL = server.URL
	_, err = Discover()
	if err == nil {
		t.Fatal("error should not be nil")
	}
	portalURL = urlBackup

	urlBackup = portalURL
	portalURL = "http://www.ninneman.org"
	_, err = Discover()
	if err == nil {
		t.Fatal("error should not be nil")
	}
	portalURL = urlBackup

	_, err = Discover()
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetState(t *testing.T) {

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

	_, err = br.Bridge.NewRequest("GET", "/", nil).Do()
	if err == nil {
		t.Fatal("should error on empty request")
	}
}
