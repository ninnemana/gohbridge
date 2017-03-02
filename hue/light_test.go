package hue

import "testing"

func TestGetLights(t *testing.T) {

	_, err := GetLights(Bridge{})
	if err == nil {
		t.Fatal("empty bridge should fail")
	}

	ll, err := Discover()
	if err != nil {
		t.Fatal(err)
	}

	if len(ll) == 0 {
		t.Log("no bridges found")
		return
	}

	_, err = GetLights(ll[0])
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetNewLights(t *testing.T) {

	_, err := GetNewLights(Bridge{})
	if err == nil {
		t.Fatal("empty bridge should fail")
	}

	ll, err := Discover()
	if err != nil {
		t.Fatal(err)
	}

	if len(ll) == 0 {
		t.Log("no bridges found")
		return
	}

	_, err = GetNewLights(ll[0])
	if err != nil {
		t.Fatal(err)
	}
}

func TestInitLightSearch(t *testing.T) {

	_, err := InitLightSearch(Bridge{}, nil)
	if err == nil {
		t.Fatal("empty bridge should fail")
	}

	ll, err := Discover()
	if err != nil {
		t.Fatal(err)
	}

	if len(ll) == 0 {
		t.Log("no bridges found")
		return
	}

	_, err = InitLightSearch(ll[0], nil)
	if err != nil {
		t.Fatal(err)
	}

	_, err = InitLightSearch(ll[0], []string{"adfjka"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetLight(t *testing.T) {

	_, err := GetLight(Bridge{}, "")
	if err == nil {
		t.Fatal("empty bridge should fail")
	}

	ll, err := Discover()
	if err != nil {
		t.Fatal(err)
	}

	if len(ll) == 0 {
		t.Log("no bridges found")
		return
	}

	lights, err := GetLights(ll[0])
	if err != nil {
		t.Fatal(err)
	}

	for id := range lights {
		_, err = GetLight(ll[0], id)
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestRenameLight(t *testing.T) {

	err := RenameLight(Bridge{}, "", "")
	if err == nil {
		t.Fatal("hould fail on empty id")
	}

	err = RenameLight(Bridge{}, "1", "")
	if err == nil {
		t.Fatal("should fail on empty name")
	}

	err = RenameLight(Bridge{}, "1", "name")
	if err == nil {
		t.Fatal("empty bridge should fail")
	}

	ll, err := Discover()
	if err != nil {
		t.Fatal(err)
	}

	if len(ll) == 0 {
		t.Log("no bridges found")
		return
	}

	lights, err := GetLights(ll[0])
	if err != nil {
		t.Fatal(err)
	}

	// err = RenameLight(ll[0], "alsdkfj", "dumby light")
	// t.Log(err)
	// if err == nil {
	// 	t.Fatal("shouldn't be able to rename light with bad id")
	// }

	for id, l := range lights {
		err = RenameLight(ll[0], id, l.Name)
		if err != nil {
			t.Fatal(err)
		}
	}
}
