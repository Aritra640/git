package main

import (
	"fmt"
	"os"

	gitcommands "github.com/Aritra640/git/app/git/git_commands"
)

func main() {

	//TODO: create a git help thing
	if len(os.Args) < 2 {
		fmt.Println("Invalid request")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		gitcommands.Gitinit()

	default:
		fmt.Println("Unknown command: ", os.Args[1])
		os.Exit(1)
	}
}
