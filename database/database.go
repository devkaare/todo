package database

// TODO: Add service utils

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"
//
//	"github.com/jackc/pgx/v5/pgxpool"
//	// _ "github.com/jackc/pgx/v5/stdlib"
// )

// type Service interface {
//	Health() map[string]string
//	Close() error
// }

// type service struct {
// 	db *pgxpool.Pool
// }
//
// var (
// 	database = os.Getenv("DB_DATABASE")
// 	password = os.Getenv("DB_PASSWORD")
// 	username = os.Getenv("DB_USERNAME")
// 	port     = os.Getenv("DB_PORT")
// 	host     = os.Getenv("DB_HOST")
// )
//
// func New() Service {
// 	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", username, password, host, port, database)
// 	db, err := pgxpool.New(context.Background(), connStr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// defer db.Close()
//
// 	return &service{
// 		db: db,
// 	}
// }
//
// func (s *service) Health() map[string]string
//
// func (s *service) Close() error
//
// func (s *service) Greeting() string {
// 	var greeting string
// 	if err := s.db.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	return greeting
// }
