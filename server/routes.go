package server

import (
	"net/http"

	"github.com/devkaare/todo/handler"
	"github.com/devkaare/todo/repository/todo"

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

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/todos/login", http.StatusSeeOther)
	})

	r.Route("/todos", s.RegisterTodoRoutes)

	return r
}

func (s *Server) RegisterTodoRoutes(r chi.Router) {
	todoHandler := &handler.Todo{
		Repo: &todo.PostgresRepo{
			Client: s.db,
		},
	}

	r.Get("/health", todoHandler.Health)

	r.Get("/login", todoHandler.GetLoginPage)
	r.Post("/login", todoHandler.LoginHandler)

	r.Group(func(r chi.Router) {
		r.Use(AuthMiddleware)

		r.Get("/", todoHandler.List)
		r.Get("/{ID}", todoHandler.GetByID)
		r.Get("/edit/{ID}", todoHandler.EditByID)

		r.Post("/", todoHandler.Create)
		r.Put("/{ID}", todoHandler.UpdateByID)
		r.Delete("/{ID}", todoHandler.DeleteByID)
	})
}
