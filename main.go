package main

import (
	"fmt"
	"log"

	"github.com/aidanmatchette/cli-todo/handlers"
)

func main() {

	var command string
	fmt.Scan(&command)

	err := handlers.ValidateInputCommand(command)
	if err != nil {
		log.Println(err)
	}

	is_created, fileErr := handlers.CheckForJsonFile()

	if fileErr != nil {
		log.Println("No JSON file found, generating file")
	}

	if !is_created {
		handlers.GenerateJsonFile()
	}
}
