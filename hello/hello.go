package main

import (
	"fmt"
	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Kenta")
	fmt.Println(message)
}
