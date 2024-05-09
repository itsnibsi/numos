// main.go

package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/itsnibsi/numos/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Welcome to Numos!\n", user.Username)
	fmt.Printf("Input your commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
