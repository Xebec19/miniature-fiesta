package main

import (
	"fmt"
	"log"

	"github.com/Xebec19/miniature-fiesta/greetings"
)

func main() {
	log.SetPrefix("greetings:")
	log.SetFlags(0)
	names := []string{"Gladys", "Samantha", "Darring"}
	message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
