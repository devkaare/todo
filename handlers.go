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
	ID          int
	Title       string
	Description string
}

var todos []Todo = []Todo{
	{ID: 1, Title: "Foo", Description: "Bar"},
	{ID: 2, Title: "Foo", Description: "Bar"},
}

func getTodoByID(ID int) (Todo, bool) {
	for i, v := range todos {
		if ID == v.ID {
			return todos[i], true
		}
	}
	return Todo{}, false
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	page := viewTodos(todos)
	page.Render(context.Background(), w)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "ID")
	ID, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Failed to read param")
		return
	}

	todo, ok := getTodoByID(ID)
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d does not exist", ID)
		return
	}

	page := viewTodo(todo)
	page.Render(context.Background(), w)
}

func getTodosHandler(w http.ResponseWriter, r *http.Request) {
	todosResult, err := json.Marshal(todos)
	if err != nil {
		w.WriteHeader(500)
		// fmt.Fprintln(w, "Failed to marshal todos")
		// return
		panic("Failed to marshal Todos")
	}
	w.Write(todosResult)
}

func uploadTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&todo); err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Failed to decode Todo")
		return
	}

	// Generate random ID if it was not set
	if todo.ID <= 0 {
		todo.ID = rand.IntN(math.MaxInt)
	}

	// Check if ID exists
	if _, ok := getTodoByID(todo.ID); ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d already exists\n", todo.ID)
		return
	}

	// Add todo to todos
	todos = append(todos, todo)

	w.Write([]byte("Successfully received post request"))
}
