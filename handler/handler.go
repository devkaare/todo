package handler

import (
	"context"
	"encoding/json"
	"log"
	"math/rand/v2"
	"net/http"

	"github.com/devkaare/todo/model"
	"github.com/devkaare/todo/repository/todo"
	"github.com/devkaare/todo/views"
)

type Todo struct {
	Repo *todo.PostgresRepo
}

func (t *Todo) Health(w http.ResponseWriter, r *http.Request) {
	jsonResp, _ := json.Marshal(t.Repo.Health())
	_, _ = w.Write(jsonResp)
}

func (t *Todo) Create(w http.ResponseWriter, r *http.Request) {
	todo := &model.Todo{}

	todo.Title = r.FormValue("title")
	todo.Description = r.FormValue("description")
	todo.ID = rand.Int()

	exists, err := t.Repo.GetTodoByID(todo.ID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exists.ID >= 0 {
		log.Println(err)
		w.WriteHeader(http.StatusConflict)
		return
	}

	if err := t.Repo.CreateTodo(todo); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// content := views.TodoList(todo)
	// content.Render(context.Background(), w)
	views.TodoPost(todo).Render(context.Background(), w)
}

func (t *Todo) List(w http.ResponseWriter, r *http.Request) {
	todos, err := t.Repo.GetTodoList()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	views.TodoForm(todos).Render(context.Background(), w)
}

func (t *Todo) GetByID(w http.ResponseWriter, r *http.Request) {
	// param := chi.URLParam(r, "ID")
	// id, err := strconv.Atoi(param)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintln(w, "failed to read id parameter")
	// 	return
	// }
	//
	// todo, ok := getTodoByID(id)
	// if !ok {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "todo with id: %d does not exist\n", id)
	// 	return
	// }
	//
	// page := views.TodoByIDConstructor(todo)
	// page.Render(context.Background(), w)
}

func (t *Todo) UpdateByID(w http.ResponseWriter, r *http.Request) {
	// param := chi.URLParam(r, "ID")
	// id, err := strconv.Atoi(param)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintln(w, "failed to read id parameter")
	// 	return
	// }
	//
	// // Check if ID exists
	// if _, ok := getTodoByID(id); !ok {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "todo with id: %d doesn't exist\n", id)
	// 	return
	// }
	//
	// todo := &database.Todo{}
	// todo.ID = id
	// todo.Title = r.FormValue("title")
	// todo.Description = r.FormValue("description")
	//
	// // Save updated todo
	// updateTodo(todo)
	//
	// fmt.Fprintln(w, "<p>Successfully updated todo!</p>")
}

func (t *Todo) DeleteByID(w http.ResponseWriter, r *http.Request) {
	// param := chi.URLParam(r, "ID")
	// id, err := strconv.Atoi(param)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintln(w, "failed to read id parameter")
	// 	return
	// }
	//
	// // Check if ID exists
	// if _, ok := getTodoByID(id); !ok {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "todo with id: %d doesn't exist\n", id)
	// 	return
	// }
	//
	// if isDeleted := deleteTodoByID(id); !isDeleted {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintln(w, "internal server error")
	// }
	//
	// fmt.Fprintln(w, "<p>Successfully deleted todo!</p>")
}

func (t *Todo) Edit(w http.ResponseWriter, r *http.Request) {
	// param := chi.URLParam(r, "ID")
	// id, err := strconv.Atoi(param)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintln(w, "failed to read id parameter")
	// 	return
	// }
	//
	// todo, ok := getTodoByID(id)
	// if !ok {
	// 	w.WriteHeader(400)
	// 	fmt.Fprintf(w, "todo with id: %d doesn't exist\n", id)
	// 	return
	// }
	//
	// content := views.TodoByIDForm(todo)
	// content.Render(context.Background(), w)
}
