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

var todoList []Todo = []Todo{
	{ID: 9223372036854775807, Title: "Example title", Description: "Example description"},
}

func getTodoByID(ID int) (Todo, bool) {
	for i, v := range todoList {
		if ID == v.ID {
			return todoList[i], true
		}
	}
	return Todo{}, false
}

func deleteTodoByID(s []Todo, i int) []Todo {
	return append(s[:i], s[i+1:]...)
}

func todoListHandler(w http.ResponseWriter, r *http.Request) {
	page := todoListComponent(todoList)
	page.Render(context.Background(), w)
}

func todoHandler(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "ID")
	ID, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Failed to read ID parameter")
		return
	}

	todo, ok := getTodoByID(ID)
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d does not exist\n", ID)
		return
	}

	page := todoComponent(todo)
	page.Render(context.Background(), w)
}

func getTodoListHandler(w http.ResponseWriter, r *http.Request) {
	todoListResult, err := json.Marshal(todoList)
	if err != nil {
		w.WriteHeader(500)
		// fmt.Fprintln(w, "Failed to marshal todoList")
		// return
		panic("Failed to marshal todoList")
	}
	w.Write(todoListResult)
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

	// Add todo to todoList
	todoList = append(todoList, todo)

	w.Write([]byte("Successfully uploaded todo!"))
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo

	todo.Title = r.FormValue("title")
	todo.Description = r.FormValue("description")

	// Generate random ID
	todo.ID = rand.IntN(math.MaxInt)

	// Check if ID exists
	if _, ok := getTodoByID(todo.ID); ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d already exists\n", todo.ID)
		return
	}

	// Add todo to todoList
	todoList = append(todoList, todo)

	result := fmt.Sprintf("<li><a href=\"/%[1]d\">%s#%[1]d</a></li>", todo.ID, todo.Title)

	fmt.Fprintln(w, result)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	param := r.FormValue("ID")
	ID, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Failed to read ID parameter")
		return
	}

	// Check if ID exists
	if _, ok := getTodoByID(ID); !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d doesn't exist\n", ID)
		return
	}

	var todo Todo

	todo.ID = ID
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
	param := r.FormValue("ID")
	ID, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Failed to read ID parameter")
		return
	}

	// Check if ID exists
	if _, ok := getTodoByID(ID); !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d doesn't exist\n", ID)
		return
	}

	// Save updated todo
	for i, v := range todoList {
		if ID == v.ID {
			// Delete slice at i (index) from todoList
			todoList = deleteTodoByID(todoList, i)
		}
	}

	result := fmt.Sprintln("<p>Successfully deleted todo!</p>")

	fmt.Fprintln(w, result)
}

func editHandler(w http.ResponseWriter, r *http.Request) {
	param := r.FormValue("ID")
	ID, err := strconv.Atoi(param)
	if err != nil {
		w.WriteHeader(400)
		fmt.Fprintln(w, "Failed to read ID parameter")
		return
	}

	// Get todo
	todo, ok := getTodoByID(ID)
	if !ok {
		w.WriteHeader(400)
		fmt.Fprintf(w, "Todo with ID: %d doesn't exist\n", ID)
		return
	}

	page := editComponent(todo)
	page.Render(context.Background(), w)
}
