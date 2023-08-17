package main

import (
	"fmt"

	"example.com/greetings"

)


func main() {
	//Get a greeting message and print
	message := greetings.Hello("Erik")
	fmt.Println(message)
}

