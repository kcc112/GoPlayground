package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Kamil")
	fmt.Printf(message)
}
