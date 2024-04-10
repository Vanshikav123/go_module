package main

import (
	"fmt"
	"log"

	"github.com/Vanshikav123/greetings"
)

func main() {

	log.SetPrefix("greetings : ")
	log.SetFlags(0)

	names := []string{
		"vanshika",
		"Aman",
		"Jyotsna",
	}

	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
