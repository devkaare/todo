package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/devkaare/todo/database"

	"github.com/go-chi/chi/v5"
)

func GetTodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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

	encoder := json.NewEncoder(w)
	if err = encoder.Encode(todo); err != nil {
		w.WriteHeader(500)
		panic("failed to encode todo")
	}
}

func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	err := encoder.Encode(todos)
	if err != nil {
		w.WriteHeader(500)
		panic("failed to encode todolist")
	}
}

func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo database.Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "failed to decode todo")
		return
	}

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
}

func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
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

	var todo database.Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "failed to decode todo")
		return
	}

	// Save updated todo
	updateTodo(todo)
}

func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
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
}
