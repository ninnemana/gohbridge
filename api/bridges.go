package api

import (
	"net/http"

	"github.com/ninnemana/gohbridge/hue"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

func ListBridges(w http.ResponseWriter, r *http.Request) {

	s, ok := r.Context().Value(contextKeyToken).(Service)
	if !ok {
		http.Error(w, "missed middleware", http.StatusInternalServerError)
		return
	}
	s.RequestHeaders = r.Header

	bridges, err := hue.Discover()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, bridges)
}

func GetBridge(w http.ResponseWriter, r *http.Request) {

	s, ok := r.Context().Value(contextKeyToken).(Service)
	if !ok {
		http.Error(w, "missed middleware", http.StatusInternalServerError)
		return
	}
	s.RequestHeaders = r.Header

	bridgeID := chi.URLParam(r, "bridgeID")

	bridges, err := hue.Discover()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, b := range bridges {
		if b.ID != bridgeID {
			continue
		}

		render.JSON(w, r, b)
		return
	}

	http.Error(w, "no bridge found", http.StatusInternalServerError)
}
