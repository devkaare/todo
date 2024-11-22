package server

import (
	"github.com/devkaare/todo/handlers"
	"github.com/go-chi/chi/v5"
)

func registerRoutes(r chi.Router) {
	r.Get("/api/v1", handlers.GetTodoListHandler)
	r.Get("/api/v1/{ID}", handlers.GetTodoHandler)
	r.Post("/api/v1/create", handlers.CreateTodoHandler)
	r.Patch("/api/v1/update/{ID}", handlers.UpdateTodoHandler)
	r.Delete("/api/v1/delete/{ID}", handlers.DeleteTodoHandler)

	r.Get("/", handlers.TodoListHandler)
	r.Get("/{ID}", handlers.TodoHandler)
	r.Post("/api/v2/create", handlers.CreateHandler)
	r.Patch("/api/v2/update/{ID}", handlers.UpdateHandler)
	r.Delete("/api/v2/delete/{ID}", handlers.DeleteHandler)
}
