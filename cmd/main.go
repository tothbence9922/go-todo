package main

import (
	"fmt"
	"gotodo/internal/hello"
)

func main() {

	greeting, err := hello.Hello("World!")

	if err != nil {
		panic("Error while constructing the greeting string!")
	}

	fmt.Println(greeting)
}
