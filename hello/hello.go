package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	//Get a greeting message and print
	message := greetings.Hello("Welcome Back Commander")
	fmt.Println(message)
}
