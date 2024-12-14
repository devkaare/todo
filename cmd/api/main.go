package main

import (
	"log"

	// "github.com/devkaare/todo/database"
	"github.com/devkaare/todo/server"
)

func main() {
	server := server.NewServer()

	// TODO: Add database

	log.Fatal(server.ListenAndServe())
}
