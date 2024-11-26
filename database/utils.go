package database

import (
	"context"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func createTable(dbInstance service) error {
	_, err := dbInstance.DB.Exec(context.Background(), "create table if not exists todos (id integer primary key, title text, description text)")
	if err != nil {
		return err
	}

	return nil
}

// func getTodos(dbInstance service) ([]Todo, error)
//
// func createTodo(todo Todo, dbInstance service) (bool, error)
//
// func getTodoByID(id int, dbInstance service) (Todo, error)
//
// func updateTodoByID(id int, dbInstance service) (bool, error)
//
// func deleteTodoByID(id int, dbInstance service) (bool, error)
