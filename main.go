package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", todoListHandler)
	r.Get("/{ID}", todoHandler)

	// TODO: Add undefined handlers
	r.Post("/api/v1/all", nil)
	r.Post("/api/v1/upload", uploadTodoHandler)
	r.Post("/api/v1/update", nil)
	r.Post("/api/v1/delete", nil)
	r.Post("/api/v1/edit", nil)
	r.Get("/api/v1/{ID}", getTodoListHandler)

	r.Post("/api/v2/upload", uploadHandler)
	r.Post("/api/v2/update/{ID}", updateHandler)
	r.Post("/api/v2/delete/{ID}", deleteHandler)
	r.Post("/api/v2/edit/{ID}", editHandler)

	http.ListenAndServe(":3000", r)
}
