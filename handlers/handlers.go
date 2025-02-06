package handlers

import (
	"errors"
	"log"
	"os"
)

var COMMANDS [6]string = [6]string{"add", "update", "delete", "list", "mark-in-progress", "mark-done"}

func CheckForJsonFile() (bool, error) {
	_, err := os.Open("task.json")

	if err != nil {
		return false, err
	}
	return true, nil
}

func GenerateJsonFile() os.File {
	json_file, err := os.Create("task.json")

	if err != nil {
		log.Fatal(err)
	}
	return *json_file
}

func ValidateInputCommand(command string) error {
	var is_valid bool = false

	for _, cmd := range &COMMANDS {
		if cmd == command {
			is_valid = true
		}
	}

	if !is_valid {
		return errors.New("Command provided is not one of the approved commands")
	}

	return nil
}

