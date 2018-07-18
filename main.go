package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/8pockets/hi/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Hi language! 🙋 \n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
