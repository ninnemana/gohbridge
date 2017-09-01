// +build integration

package hue

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/twinj/uuid"
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
