package hue

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

type ConfigHandler struct {
	sync.Mutex
}

func (h *ConfigHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

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

func TestGetConfig(t *testing.T) {
	var b Bridge
	_, err := GetConfig(b)
	if err == nil {
		t.Fatal("should have errored on bad bridge")
	}

	bridges, err := Discover()
	if err != nil {
		t.Fatal(err)
	}

	if bridges == nil || len(bridges) == 0 {
		t.Log("no bridges found")
		return
	}

	handler := &ConfigHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	b = bridges[0]
	ip := b.InternalIP
	b.InternalIP = strings.Replace(server.URL, "http://", "", 1)
	b.User = "random_user"
	_, err = GetConfig(b)
	if err == nil {
		t.Fatal("should have errored on bad username")
	}

	b.User = testUser
	b.InternalIP = ip

	_, err = GetConfig(b)
	if err != nil {
		t.Fatal(err)
	}
}
