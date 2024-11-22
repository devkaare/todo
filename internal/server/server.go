package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var port = ":8080" // TODO: Move to .env

func New() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	registerRoutes(r)

	log.Println("Started listening on port", port)
	if err := http.ListenAndServe(port, r); err != nil {
		log.Fatal(err)
	}
}
