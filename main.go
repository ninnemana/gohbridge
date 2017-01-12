package main

import (
	"flag"
	"fmt"
	"net/http"

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

	http.ListenAndServe(":8080", r)
}
