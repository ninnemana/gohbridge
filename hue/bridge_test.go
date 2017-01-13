package hue

import "testing"

func TestDiscover(t *testing.T) {
	urlBackup := portalURL
	portalURL = "http://www.ninneman.org"
	_, err := Discover()
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

	_, err = ll[0].GetState()
	if err != nil {
		t.Fatal(err)
	}
}
