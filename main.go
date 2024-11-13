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

	r.Get("/api/v1", getTodoListHandler)
	r.Post("/api/v1", uploadTodoHandler)

	r.Post("/api/v2/upload", uploadHandler)
	r.Post("/api/v2/update", updateHandler)
	r.Post("/api/v2/delete", deleteHandler)
	r.Post("/api/v2/edit", editHandler)

	http.ListenAndServe(":3000", r)
}
