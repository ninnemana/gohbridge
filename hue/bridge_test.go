package hue

import (
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	})
	http.ListenAndServe(":9999", nil)

	m.Run()
}

func TestDiscover(t *testing.T) {

	urlBackup := portalURL
	portalURL = "asl;dfj"
	_, err := Discover()
	if err == nil {
		t.Fatal("error should not be nil")
	}
	portalURL = urlBackup

	urlBackup = portalURL
	portalURL = "http://localhost:9999"
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

	_, err = ll[0].GetState()
	if err != nil {
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
