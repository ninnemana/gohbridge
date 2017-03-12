package api

import (
	"net/http"

	"github.com/ninnemana/gohbridge/hue"
	"github.com/pressly/chi/render"
)

// GetConfig returns the configuration of the bridge.
func GetConfig(w http.ResponseWriter, r *http.Request) {

	s, ok := r.Context().Value(contextKeyToken).(Service)
	if !ok {
		http.Error(w, "missed middleware", http.StatusInternalServerError)
		return
	}
	s.RequestHeaders = r.Header

	cfg, err := hue.GetConfig(*s.Bridge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, cfg)
}
