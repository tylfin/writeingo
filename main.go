package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"monkey/repl"
)

// main is the entry point of the program
func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatalf("failed to get current user: %v", err)
	}

	fmt.Printf("Hello %s! This is the Monkey programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
