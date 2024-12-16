// TODO:
// 1. Add more tests for DB.
// 2. Update handlers to support new DB.
// 3. Update handlers to render using templ.
// 4. Update views structure.

package server

import (
	"encoding/json"
	"log"
	"net/http"

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

	r.Get("/", s.HelloWorldHandler)

	r.Get("/health", s.healthHandler)

	r.Route("/todo", func(r chi.Router) {
		r.Post("/", s.CreateHandler)
		r.Get("/", s.ListHandler)
		r.Get("/{ID}", s.GetByIDHandler)
		r.Put("/{ID}", s.UpdateByIDHandler)
		r.Delete("/{ID}", s.DeleteByIDHandler)
	})

	return r
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

func (s *Server) CreateHandler(w http.ResponseWriter, r *http.Request) {
	// todo := &database.Todo{}
	//
	// todo.Title = r.FormValue("title")
	// todo.Description = r.FormValue("description")
	//
	// // Generate unique ID
	// todo.ID = rand.IntN(math.MaxInt)
	//
	// // Check if ID exists
	// if _, ok := getTodoByID(todo.ID); ok {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "todo with id: %d already exists\n", todo.ID)
	// 	return
	// }
	//
	// // Add todo to todoList
	// todos = append(todos, todo)
	//
	// content := views.TodoList(todo)
	// content.Render(context.Background(), w)
}

func (s *Server) ListHandler(w http.ResponseWriter, r *http.Request) {
	// page := views.TodosConstructor(todos)
	// page.Render(context.Background(), w)
}

func (s *Server) GetByIDHandler(w http.ResponseWriter, r *http.Request) {
	// param := chi.URLParam(r, "ID")
	// id, err := strconv.Atoi(param)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintln(w, "failed to read id parameter")
	// 	return
	// }
	//
	// todo, ok := getTodoByID(id)
	// if !ok {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "todo with id: %d does not exist\n", id)
	// 	return
	// }
	//
	// page := views.TodoByIDConstructor(todo)
	// page.Render(context.Background(), w)
}

func (s *Server) UpdateByIDHandler(w http.ResponseWriter, r *http.Request) {
	// param := chi.URLParam(r, "ID")
	// id, err := strconv.Atoi(param)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintln(w, "failed to read id parameter")
	// 	return
	// }
	//
	// // Check if ID exists
	// if _, ok := getTodoByID(id); !ok {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "todo with id: %d doesn't exist\n", id)
	// 	return
	// }
	//
	// todo := &database.Todo{}
	// todo.ID = id
	// todo.Title = r.FormValue("title")
	// todo.Description = r.FormValue("description")
	//
	// // Save updated todo
	// updateTodo(todo)
	//
	// fmt.Fprintln(w, "<p>Successfully updated todo!</p>")
}

func (s *Server) DeleteByIDHandler(w http.ResponseWriter, r *http.Request) {
	// param := chi.URLParam(r, "ID")
	// id, err := strconv.Atoi(param)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintln(w, "failed to read id parameter")
	// 	return
	// }
	//
	// // Check if ID exists
	// if _, ok := getTodoByID(id); !ok {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "todo with id: %d doesn't exist\n", id)
	// 	return
	// }
	//
	// if isDeleted := deleteTodoByID(id); !isDeleted {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintln(w, "internal server error")
	// }
	//
	// fmt.Fprintln(w, "<p>Successfully deleted todo!</p>")
}

func (s *Server) EditHandler(w http.ResponseWriter, r *http.Request) {
	// param := chi.URLParam(r, "ID")
	// id, err := strconv.Atoi(param)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintln(w, "failed to read id parameter")
	// 	return
	// }
	//
	// todo, ok := getTodoByID(id)
	// if !ok {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "todo with id: %d doesn't exist\n", id)
	// 	return
	// }
	//
	// content := views.TodoByIDForm(todo)
	// content.Render(context.Background(), w)
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(s.db.Health())
	_, _ = w.Write(jsonResp)
}
