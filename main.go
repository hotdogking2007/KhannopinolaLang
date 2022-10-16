package main

import (
	"KhannopinolaLang/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello, %s", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
