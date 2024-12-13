package database

import (
	// "context"

	// "github.com/jackc/pgx/v5"
	_ "github.com/joho/godotenv/autoload"
)

type Todo struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// func CreateTable(dbInstance service) error {
// 	_, err := dbInstance.DB.Exec(context.Background(), "create table if not exists todos (id integer primary key, title text, description text)")
// 	if err != nil {
// 		return err
// 	}
//
// 	return nil
// }

// TODO: Add these:
// GetTodos(db *pgx.Conn) ([]Todo, error)
// CreateTodo(todo Todo, db *pgx.Conn) (bool, error)
// GetTodoByID(id int, db *pgx.Conn) (Todo, error)
// UpdateTodoByID(id int, db *pgx.Conn) (bool, error)
// DeleteTodoByID(id int, db *pgx.Conn) (bool, error)
