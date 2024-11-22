package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/devkaare/todo/internal/database"
	"github.com/devkaare/todo/views"

	"github.com/go-chi/chi/v5"
)

var todoList []database.Todo = []database.Todo{
	{
		ID:          9223372036854775807,
		Title:       "example title",
		Description: "example description",
	},
}

func getTodoByID(id int) (database.Todo, bool) {
	for i, v := range todoList {
		if v.ID == id {
			return todoList[i], true
		}
	}
	return database.Todo{}, false
}

// API Handlers (V1)
func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
		fmt.Fprintf(w, "Todo with ID: %d doesn't exist\n", id)
		return
	}

	encoder := json.NewEncoder(w)
	if err = encoder.Encode(todo); err != nil {
		w.WriteHeader(500)
		panic("Failed to encode todo")
	}
}

func GetTodoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	err := encoder.Encode(todoList)
	if err != nil {
		w.WriteHeader(500)
		panic("Failed to encode todoList")
	}
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo database.Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Failed to decode Todo")
		return
	}

	// Generate unique ID
	todo.ID = rand.IntN(math.MaxInt)

	// Check if ID exists
	if _, ok := getTodoByID(todo.ID); ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d already exists\n", todo.ID)
		return
	}

	// Add todo to todoList
	todoList = append(todoList, todo)
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
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
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Failed to decode Todo")
		return
	}

	// Save updated todo
	for i, v := range todoList {
		if v.ID == id {
			todo.ID = id // Update todo.ID because the code below replaces the ENTIRE todo
			todoList[i] = todo
		}
	}
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
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

	// Save updated todo
	for i, v := range todoList {
		if v.ID == id {
			// Delete todo at index from todoList
			todoList = append(todoList[:i], todoList[i+1:]...)
		}
	}
}

// API Handlers (V2)
func TodoListHandler(w http.ResponseWriter, r *http.Request) {
	page := views.TodosConstructor(todoList)
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
	todoList = append(todoList, todo)

	fmt.Fprintln(w, fmt.Sprintf("<li><a href=\"/%[1]d\">%s#%[1]d</a></li>", todo.ID, todo.Title))
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
	for i, v := range todoList {
		if todo.ID == v.ID {
			todoList[i] = todo
		}
	}

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

	// Save updated todo
	for i, v := range todoList {
		if id == v.ID {
			// Delete todo at index from todoList
			todoList = append(todoList[:i], todoList[i+1:]...)
		}
	}

	fmt.Fprintln(w, "<p>Successfully deleted todo!</p>")
}
