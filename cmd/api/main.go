package main

import (
	"log"

	// "github.com/devkaare/todo/database"
	"github.com/devkaare/todo/server"
)

func main() {
	// TODO: Add proper database stuff here and pass to server
	server := server.NewServer()

	// log.Println("Connecting to database...")
	// dbInstance := database.New()

	// greeting := database.Greeting(dbInstance)
	// log.Println(greeting)

	log.Fatal(server.ListenAndServe())
}
