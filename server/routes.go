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
	r.Get("/todo/{ID}", handlers.TodoByIDHandler)
	r.Get("/todo/edit/{ID}", handlers.EditHandler)
	r.Post("/todo/create", handlers.CreateHandler)
	r.Patch("/todo/update/{ID}", handlers.UpdateHandler)
	r.Delete("/todo/delete/{ID}", handlers.DeleteHandler)

	return r
}
