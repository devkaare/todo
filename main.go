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

	r.Post("/api/v2/upload", uploadHandler)
	r.Get("/api/v2/update/form/{ID}", todoUpdateHandler)
	r.Patch("/api/v2/update/{ID}", updateHandler)
	r.Delete("/api/v2/delete/{ID}", deleteHandler)

	r.Get("/api/v1/", getTodoListHandler)
	r.Get("/api/v1/{ID}", getTodoHandler)
	r.Post("/api/v1/upload", uploadTodoHandler)
	r.Patch("/api/v1/update", updateTodoHandler)
	r.Delete("/api/v1/delete", nil)

	http.ListenAndServe(":3000", r)
}
