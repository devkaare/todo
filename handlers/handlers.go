package handlers

import (
	"context"
	"fmt"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/devkaare/todo/database"
	"github.com/devkaare/todo/views"

	"github.com/go-chi/chi/v5"
)

var todos []*database.Todo

func getTodoByID(id int) (*database.Todo, bool) {
	for i, todoFromTodos := range todos {
		if todoFromTodos.ID == id {
			// Return todo at current index in todos
			return todos[i], true
		}
	}
	return &database.Todo{}, false
}

func deleteTodoByID(id int) bool {
	for i, todoFromTodos := range todos {
		if todoFromTodos.ID == id {
			// Delete todo at current index in todos
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}

func updateTodo(todo *database.Todo) bool {
	for i, todoFromTodos := range todos {
		if todoFromTodos.ID == todo.ID {
			// Update todo at current index in todos
			todos[i] = todo
			return true
		}
	}
	return false
}

func TodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "ID")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "failed to read id parameter")
		return
	}

	todo, ok := getTodoByID(id)
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "todo with id: %d does not exist\n", id)
		return
	}

	page := views.TodoByIDConstructor(todo)
	page.Render(context.Background(), w)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	todo := &database.Todo{}

	todo.Title = r.FormValue("title")
	todo.Description = r.FormValue("description")

	// Generate unique ID
	todo.ID = rand.IntN(math.MaxInt)

	// Check if ID exists
	if _, ok := getTodoByID(todo.ID); ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "todo with id: %d already exists\n", todo.ID)
		return
	}

	// Add todo to todoList
	todos = append(todos, todo)

	content := views.TodoList(todo)
	content.Render(context.Background(), w)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "ID")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "failed to read id parameter")
		return
	}

	// Check if ID exists
	if _, ok := getTodoByID(id); !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "todo with id: %d doesn't exist\n", id)
		return
	}

	todo := &database.Todo{}
	todo.ID = id
	todo.Title = r.FormValue("title")
	todo.Description = r.FormValue("description")

	// Save updated todo
	updateTodo(todo)

	fmt.Fprintln(w, "<p>Successfully updated todo!</p>")
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "ID")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "failed to read id parameter")
		return
	}

	// Check if ID exists
	if _, ok := getTodoByID(id); !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "todo with id: %d doesn't exist\n", id)
		return
	}

	if isDeleted := deleteTodoByID(id); !isDeleted {
		w.WriteHeader(400)
		fmt.Fprintln(w, "internal server error")
	}

	fmt.Fprintln(w, "<p>Successfully deleted todo!</p>")
}

func EditHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "ID")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "failed to read id parameter")
		return
	}

	todo, ok := getTodoByID(id)
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "todo with id: %d doesn't exist\n", id)
		return
	}

	content := views.TodoByIDForm(todo)
	content.Render(context.Background(), w)
}

func TodosHandler(w http.ResponseWriter, r *http.Request) {
	page := views.TodosConstructor(todos)
	page.Render(context.Background(), w)
}
