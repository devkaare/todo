package main

import (
	"encoding/json"
	"fmt"
	"math"
	"math/rand/v2"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Todo struct {
	ID          int
	Title       string
	Description string
}

var todos []Todo = []Todo{
	{ID: 1, Title: "Foo", Description: "Bar"},
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		todosResult, err := json.Marshal(todos)
		if err != nil {
			w.WriteHeader(500)
			// fmt.Fprintln(w, "Failed to marshal todos")
			// return
			panic("Failed to marshal todos")
		}
		w.Write(todosResult)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var todo Todo
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&todo); err != nil {
			w.WriteHeader(400)
			fmt.Fprintln(w, "Failed to decode todo")
			return
		}

		// Generate random ID if it was not set
		if todo.ID <= 0 {
			todo.ID = rand.IntN(math.MaxInt8)
		}

		// Check if ID exists
		for _, v := range todos {
			if v.ID == todo.ID {
				w.WriteHeader(400)
				fmt.Fprintf(w, "Todo with ID: %d already exists\n", todo.ID)
				return
			}
		}

		// Add todo to todos
		todos = append(todos, todo)

		w.Write([]byte("Successfully received post request"))
	})

	http.ListenAndServe(":3000", r)
}
