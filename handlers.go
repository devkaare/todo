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

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")

	// TODO: Add todo to todos here
	fmt.Println(title, description)

	// <li><a href={ templ.URL(strconv.Itoa(todo.ID)) }>{ todo.Title }#{ strconv.Itoa(todo.ID) }</a></li>
	entry := fmt.Sprintf("<li><a href=\"/%[1]d\">%s#%[1]d</a></li>", 69420, title)

	fmt.Fprintln(w, entry)
}

func updateHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	description := r.FormValue("description")

	fmt.Println(title, description)
	fmt.Println(r.Method)

	// <header>
	// 	<h1>{ todo.Title }#{ strconv.Itoa(todo.ID) }</h1>
	// </header>
	// <body>
	// 	<p>{ todo.Description }</p>
	// </body>
	todo := fmt.Sprintf(`
		<header>
			<h1>%s#%d</h1>
		</header>
		<body>
			<p>%s<p>
		</body>
		`, title, 69420, description)

	fmt.Fprintln(w, todo)
}
