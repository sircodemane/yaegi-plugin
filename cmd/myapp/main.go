package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/sircodemane/yaegi-plugin/internal"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", internal.HelloWorldHandler)
	r.Get("/greet", internal.GreetHandler)

	if plugin, err := internal.LoadPlugin(); err != nil {
		log.Fatalf("failed to load plugin: %s", err)
	} else {
		r.Route("/plugin", func(r chi.Router) {
			for pattern, handler := range plugin.Routes() {
				r.Get(pattern, handler)
			}
		})
	}

	fmt.Println("server running...")
	http.ListenAndServe(":3000", r)
}
