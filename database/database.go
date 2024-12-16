package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Service interface {
	Health() map[string]string

	Close() error
}

type service struct {
	db *sql.DB
}

type Todo struct {
	ID          int
	Title       string
	Description string
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *service
)

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}
	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err)
		return stats
	}

	stats["status"] = "up"
	stats["message"] = "It's healthy"

	return stats
}

func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", database)
	return s.db.Close()
}

func (s *service) CreateTodo(todo *Todo) error {
	_, err := s.db.Exec(
		"INSERT INTO todo (id, title, description) VALUES (?, ?, ?)",
		todo.ID, todo.Title, todo.Description,
	)
	if err != nil {
		return fmt.Errorf("CreateTodo: %v", err)
	}

	// TODO: Check for existing todo with ID

	return nil
}

func (s *service) GetTodoList() ([]Todo, error) {
	var todos []Todo

	rows, err := s.db.Query("SELECT * FROM todo")
	if err != nil {
		return todos, err
	}
	defer rows.Close()

	for rows.Next() {
		var todo Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
			return nil, fmt.Errorf("GetTodoList %d: %v", todo.ID, err)
		}
		todos = append(todos, todo)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetTodoList %v:", err)
	}
	return todos, nil
}

func (s *service) GetTodoByID(id int) (Todo, error) {
	var todo Todo

	row := s.db.QueryRow("SELECT * FROM todo WHERE id = ?")
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Description); err != nil {
		if err == sql.ErrNoRows {
			return todo, fmt.Errorf("GetTodoByID %d: no such todo", id)
		}
		return todo, fmt.Errorf("GetTodoByID %d: %v", id, err)
	}
	return todo, nil

}

func (s *service) UpdateTodoByID(todo *Todo) error {
	_, err := s.db.Exec("UPDATE todo SET title = $2, description = $3 WHERE id = $1", todo.ID, todo.Title, todo.Description)
	if err != nil {
		return fmt.Errorf("CreateTodo: %v", err)
	}
	return nil
}

func (s *service) DeleteTodoByID(id int) error {
	result, err := s.db.Exec("DELETE FROM todo WHERE id = ?")
	if err != nil {
		return fmt.Errorf("DeleteTodoByID %d, %v", id, err)
	}
	count, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("DeleteTodoByID %d: %v", id, err)
	}
	if count < 1 {
		return fmt.Errorf("DeleteTodoByID %d: no such todo", id)
	}
	return nil
}
