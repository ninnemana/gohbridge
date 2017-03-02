package api

import (
	"net/http"

	"github.com/ninnemana/gohbridge/hue"
	"github.com/pressly/chi"
	"github.com/pressly/chi/render"
)

func ListLights(w http.ResponseWriter, r *http.Request) {

	s, ok := r.Context().Value(contextKeyToken).(Service)
	if !ok {
		http.Error(w, "missed middleware", http.StatusInternalServerError)
		return
	}
	s.RequestHeaders = r.Header

	lights, err := hue.GetLights(*s.Bridge)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, lights)
}

func GetLight(w http.ResponseWriter, r *http.Request) {

	s, ok := r.Context().Value(contextKeyToken).(Service)
	if !ok {
		http.Error(w, "missed middleware", http.StatusInternalServerError)
		return
	}
	s.RequestHeaders = r.Header

	lightID := chi.URLParam(r, "lightID")

	light, err := hue.GetLight(*s.Bridge, lightID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	render.JSON(w, r, light)
}
