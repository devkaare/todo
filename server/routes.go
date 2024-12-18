// TODO:
// 2. Update handlers to support new DB.
// 3. Update handlers to render using templ.
// 4. Update views structure.

package server

import (
	"encoding/json"
	"log"
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

	// r.Get("/", s.HelloWorldHandler)

	// r.Get("/", templ.Handler(views.HelloForm()).ServeHTTP)
	// r.Post("/hello", views.HelloWebHandler)

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
	r.Post("/", todoHandler.Create)
	r.Get("/", todoHandler.List)
	r.Get("/{ID}", todoHandler.GetByID)
	r.Put("/{ID}", todoHandler.UpdateByID)
	r.Delete("/{ID}", todoHandler.DeleteByID)
}

func (s *Server) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
