package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type service struct {
	DB *pgxpool.Pool
}

var (
	database = os.Getenv("DB_DATABASE")
	password = os.Getenv("DB_PASSWORD")
	username = os.Getenv("DB_USERNAME")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
)

func New() service {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
	dbpool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	// defer dbpool.Close()

	return service{
		DB: dbpool,
	}
}

func Greeting(dbpool service) string {
	var greeting string
	if err := dbpool.DB.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting); err != nil {
		log.Fatal(err)
	}

	return greeting
}
