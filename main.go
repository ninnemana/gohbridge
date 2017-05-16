// 97.90.224.136

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/ninnemana/gohbridge/api"
	"github.com/pressly/chi"
	"github.com/pressly/chi/docgen"
	"github.com/pressly/chi/middleware"
)

var routes = flag.Bool("routes", false, "Generate router documentation")

func main() {
	flag.Parse()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)

	r.Mount("/", protectedRoutes())

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root."))
	})

	r.Get("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up"))
	})

	if *routes {
		fmt.Println(docgen.MarkdownRoutesDoc(r, docgen.MarkdownOpts{
			ProjectPath: "github.com/ninnemana/gohbridge",
			Intro:       "API Interface to interact with Hue, Alexa, Google Home",
		}))
		return
	}

	log.Println("Starting server...")
	http.ListenAndServe(":8080", r)
}

func protectedRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(api.WithService)

	r.Route("/bridge", func(r chi.Router) {
		r.Get("/", api.ListBridges)
		r.Get("/:bridgeID", api.GetBridge)
	})

	r.Route("/light", func(r chi.Router) {
		r.Get("/", api.ListLights)
		r.Get("/:lightID", api.GetLight)
	})

	r.Route("/config", func(r chi.Router) {
		r.Get("/", api.GetConfig)
	})

	return r
}
