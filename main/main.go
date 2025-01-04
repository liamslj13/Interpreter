package main

import (
	"WritingAnInterpreterInGO/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s!\n", user.Username)
	fmt.Printf("Type some commands below\n")
	repl.Start(os.Stdin, os.Stdout)
}
