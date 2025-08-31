package main

import (
	"fmt"
	"log"
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
	case "help":
		gitcommands.Githelp()

	case "cat-file": 
		gitcommands.Gitcatfile(os.Args[2] , os.Args[3])

	case "hash-object": 
	
		command := gitcommands.GitHashObject{}
		RunCommand(&command)

	case "status": 
		log.Println("show status")

	case "add": 

	case "commit": 

	case "push":

	default:

		msg := fmt.Sprintf("git: '%v' is not a git command. See 'git --help'." , os.Args[1])
		os.Stdout.Write([]byte(msg))
	}
}


func RunCommand(command gitcommands.GitClient) {
	command.Init()
	command.Execute()
}
