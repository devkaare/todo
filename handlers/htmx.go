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

func TodosHandler(w http.ResponseWriter, r *http.Request) {
	page := views.TodosConstructor(todos)
	page.Render(context.Background(), w)
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "ID")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Failed to read ID parameter")
		return
	}

	todo, ok := getTodoByID(id)
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d does not exist\n", id)
		return
	}

	page := views.TodoConstructor(todo)
	page.Render(context.Background(), w)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	var todo database.Todo

	todo.Title = r.FormValue("title")
	todo.Description = r.FormValue("description")

	// Generate unique ID
	todo.ID = rand.IntN(math.MaxInt)

	// Check if ID exists
	if _, ok := getTodoByID(todo.ID); ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d already exists\n", todo.ID)
		return
	}

	// Add todo to todoList
	todos = append(todos, todo)

	fmt.Fprintf(w, "<li><a href=\"/%[1]d\">%s#%[1]d</a></li>", todo.ID, todo.Title)
}

func UpdateHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "ID")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Failed to read ID parameter")
		return
	}

	// Check if ID exists
	if _, ok := getTodoByID(id); !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d doesn't exist\n", id)
		return
	}

	var todo database.Todo
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
		fmt.Fprintln(w, "Failed to read ID parameter")
		return
	}

	// Check if ID exists
	if _, ok := getTodoByID(id); !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d doesn't exist\n", id)
		return
	}

	deleteTodoByID(id)

	fmt.Fprintln(w, "<p>Successfully deleted todo!</p>")
}
