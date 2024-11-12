package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// NOTE: This does render updates to todos, so im using r.Get kinda like middleware instead (see handlers)
	// r.Handle("/all", templ.Handler(viewTodos(todos)))

	r.Get("/", todosHandler)
	r.Get("/{ID}", todoHandler)

	r.Get("/api/v1", getTodosHandler)
	r.Post("/api/v1", uploadTodoHandler)

	http.ListenAndServe(":3000", r)
}
