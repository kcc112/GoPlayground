package greetings

import (
	"errors"
	"fmt"
)

// Hello get a greeting message and print it
func Hello(name string) (string, error) {

	if name == "" {
		return "", errors.New("empty name")
	} 

	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
