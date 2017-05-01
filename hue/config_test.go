package hue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"sync"
	"testing"
)

type ConfigHandler struct {
	sync.Mutex

	Fail string
}

func (h *ConfigHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var obj interface{}
	switch h.Fail {
	case BadUsername:
		w.WriteHeader(500)
		return
	default:
		obj = Config{}
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

func TestCreateUser(t *testing.T) {

	tmp := portalURL

	handler := &DiscoverHandler{}
	server := httptest.NewServer(handler)
	defer server.Close()

	portalURL = server.URL
	defer func(str string) {
		portalURL = str
	}(tmp)

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
	handler := &ConfigHandler{
		Fail: BadUsername,
	}
	server := httptest.NewServer(handler)
	defer server.Close()

	configHost, err := url.Parse(server.URL)
	if err != nil {
		t.Fatal(err)
	}

	discover := &DiscoverHandler{
		Address: fmt.Sprintf("%s", configHost.Host),
	}
	discoverServer := httptest.NewServer(discover)
	defer discoverServer.Close()
	portalURL = discoverServer.URL

	var b Bridge
	_, err = GetConfig(b)
	if err == nil {
		t.Fatal("should have errored on bad bridge")
	}

	bridges, err := Discover()
	if err != nil {
		t.Fatal(err)
	}

	if bridges == nil {
		t.Log("no bridges found")
		return
	}

	portalURL = server.URL

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

	handler.Fail = ""
	server = httptest.NewServer(handler)
	defer server.Close()

	_, err = GetConfig(b)
	if err != nil {
		t.Fatal(err)
	}
}
