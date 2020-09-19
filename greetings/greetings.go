package greetings

import "fmt"

// Hello get a greeting message and print it
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
