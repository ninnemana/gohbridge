package hue

import "testing"

func TestCreateUser(t *testing.T) {

	ll, err := Discover()
	if err != nil {
		t.Fatal(err)
	}

	if len(ll) == 0 {
		t.Log("no bridges found")
		return
	}

	user, err := CreateUser(ll[0])
	t.Log(user, err)
}
