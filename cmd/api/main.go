package main

import (
	"log"

	// "github.com/devkaare/todo/database"
	"github.com/devkaare/todo/server"
)

func main() {
	// TODO: Add proper database stuff here and pass to server
	server := server.NewServer()

	// TODO: dbInstance := database.New()

	log.Fatal(server.ListenAndServe())
}
