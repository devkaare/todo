package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/api/v1", getTodoListHandler)
	r.Get("/api/v1/{ID}", getTodoHandler)
	r.Post("/api/v1/create", createTodoHandler)
	r.Patch("/api/v1/update/{ID}", updateTodoHandler)
	r.Delete("/api/v1/delete/{ID}", deleteTodoHandler)

	r.Get("/", todoListHandler)
	r.Get("/{ID}", todoHandler)
	r.Post("/api/v2/create", createHandler)
	r.Get("/api/v2/update/form/{ID}", todoUpdateHandler)
	r.Patch("/api/v2/update/{ID}", updateHandler)
	r.Delete("/api/v2/delete/{ID}", deleteHandler)

	http.ListenAndServe(":3000", r)
}
