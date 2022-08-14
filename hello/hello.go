package main

import (
	"fmt"

	"log"

	"artetech/greetings"
)

func main() {
	log.SetPrefix("Greetings: ")

	log.SetFlags(0)

	names := []string{"Jim Bo", "Bob Billy", "Carl The Bogi"}

	messages, error := greetings.Hellos(names)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println(messages)
}
