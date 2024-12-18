package views

import (
	"log"
	"net/http"
	"strconv"

	"github.com/devkaare/todo/model"
	"github.com/go-chi/chi/v5"
)

func TodoWebHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}

	todo := &model.Todo{}

	todo.ID, err = strconv.Atoi(chi.URLParam(r, "ID"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	todo.Title = r.FormValue("title")
	todo.Description = r.FormValue("description")

	component := TodoPost(todo)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in TodoWebHandler: %e", err)
	}
}
