package main

import (
	"github.com/devkaare/todo/server"
	"log"
)

func main() {
	// TODO: Add proper database stuff here and pass to server
	server := server.NewServer()

	log.Fatal(server.ListenAndServe())
}
