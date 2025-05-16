package main

// go mod init path to intialise go project(go.mod will be added to the directory)

//  go run . command runs file with main package as default

// go get       - Adds, updates, or removes specific Go modules manually.

// go mod tidy  - Automatically adds missing and removes unused dependencies in go.mod and go.sum.

import (
	"example/greetings"
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())
	greeting, greetingError := greetings.Greet("")
	if greetingError != nil {
		fmt.Println(greetingError.Error())
	} else {
		fmt.Println(greeting)
	}
}
