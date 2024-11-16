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

	r.Get("/api/v1/{ID}", getTodoListHandler) // ID
	r.Post("/api/v1/upload", uploadTodoHandler)
	r.Post("/api/v1/update", nil) // TODO: Add handlers for these
	r.Post("/api/v1/delete", nil)
	r.Post("/api/v1/edit", nil)

	r.Post("/api/v2/upload", uploadHandler)
	r.Post("/api/v2/update/{ID}", updateHandler) // ID PUT
	r.Post("/api/v2/delete/{ID}", deleteHandler) // ID DELETE
	r.Post("/api/v2/edit/{ID}", editHandler)     // ID PUT

	http.ListenAndServe(":3000", r)
}
