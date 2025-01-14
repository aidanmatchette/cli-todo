package main

import (
	"errors"
	"fmt"
	"log"
)

func validateInputCommand(command string, allowed_commands [3]string) (error) {
    var is_valid bool = false

    for _, cmd := range allowed_commands {
        if cmd == command {
            is_valid = true
        }
    }

    if !is_valid {
        return errors.New("Command provided is not one of the approved commands")
    }

    return nil
}


func main() {
    var allowed_commands [3]string = [3]string{"add", "update", "delete"}
    fmt.Println(allowed_commands)

    var command string
    fmt.Scan(&command)

    err := validateInputCommand(command, allowed_commands)
    if err != nil {
        log.Fatal(err)
    }

}
