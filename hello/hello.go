package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Emin", "Bill", "Boris"}

	messages, err := greetings.Hellos(names)

	//error handling
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
