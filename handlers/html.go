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
		fmt.Fprintln(w, "failed to read id parameter")
		return
	}

	todo, ok := getTodoByID(id)
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "todo with id: %d does not exist\n", id)
		return
	}

	page := views.TodoConstructor(todo)
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

	fmt.Fprintf(w, `<li><a href="%d">%s</a></li>`, todo.ID, todo.Title)
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

	fmt.Fprintf(w, `
			<form hx-target="this" hx-swap="outerHTML" autocomplete="off">
				<input type="text" name="title" value="%s" required>
				<textarea name="description" required>%s</textarea>
				<button type="submit" hx-patch="/api/v2/update/%d">Submit</button>
			</form>
		`,
		todo.Title, todo.Description, todo.ID)
}
