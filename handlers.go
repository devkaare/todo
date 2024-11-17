package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

var todoList []Todo = []Todo{
	{
		ID:          9223372036854775807,
		Title:       "example title",
		Description: "example description",
	},
}

func getTodoByID(id int) (Todo, bool) {
	for i, v := range todoList {
		if v.ID == id {
			return todoList[i], true
		}
	}
	return Todo{}, false
}

// API Handlers (V1)
func getTodoHandler(w http.ResponseWriter, r *http.Request) {
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

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	err := encoder.Encode(todoList)
	if err != nil {
		w.WriteHeader(500)
		panic("Failed to encode todoList")
	}
}

func createTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
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

func updateTodoHandler(w http.ResponseWriter, r *http.Request) {
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

	var todo Todo
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

func deleteTodoHandler(w http.ResponseWriter, r *http.Request) {
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
func todoListHandler(w http.ResponseWriter, r *http.Request) {
	page := todoListComponent(todoList)
	page.Render(context.Background(), w)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
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

	page := todoComponent(todo)
	page.Render(context.Background(), w)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo

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

func updateHandler(w http.ResponseWriter, r *http.Request) {
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

	var todo Todo
	todo.ID = id
	todo.Title = r.FormValue("title")
	todo.Description = r.FormValue("description")

	// Save updated todo
	for i, v := range todoList {
		if todo.ID == v.ID {
			todoList[i] = todo
		}
	}

	page := detailComponent(todo)
	page.Render(context.Background(), w)
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
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

// NOTE: This is different from updateHandler, because this
// returns the form for updateHandler requests (instead of actual updating)
func todoUpdateHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "ID")
	id, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Failed to read ID parameter")
		return
	}

	// Get todo and check if ID exists
	todo, ok := getTodoByID(id)
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d doesn't exist\n", id)
		return
	}

	page := editComponent(todo)
	page.Render(context.Background(), w)
}
