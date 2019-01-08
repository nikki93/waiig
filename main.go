package main

import (
	"fmt"
	"os"
	"os/user"
	"waiig/repl"
)

func main() {
	currentUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("hello %s -- welcome to this thing\n", currentUser.Username)
	repl.Start(os.Stdin, os.Stdout)
}
