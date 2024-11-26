package server

import (
	"net/http"

	"github.com/devkaare/todo/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	r.Get("/", handlers.TodosHandler)
	r.Get("/{ID}", handlers.TodoHandler)

	r.Get("/api/v1", handlers.GetTodosHandler)
	r.Get("/api/v1/{ID}", handlers.GetTodoHandler)
	r.Post("/api/v1/create", handlers.CreateTodoHandler)
	r.Patch("/api/v1/update/{ID}", handlers.UpdateTodoHandler)
	r.Delete("/api/v1/delete/{ID}", handlers.DeleteTodoHandler)

	r.Get("/api/v2/edit/{ID}", handlers.EditHandler)
	r.Get("/api/v2/{ID}", handlers.GetTodoHandler)
	r.Post("/api/v2/create", handlers.CreateHandler)
	r.Patch("/api/v2/update/{ID}", handlers.UpdateHandler)
	r.Delete("/api/v2/delete/{ID}", handlers.DeleteHandler)

	return r
}
