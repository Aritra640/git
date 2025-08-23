package main

import (
	"fmt"
	"os"
)

func main() {

	//TODO: create a git help thing
	if len(os.Args) < 2 {
		fmt.Println("Invalid request")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":

	default:
		fmt.Println("Unknown command: ", os.Args[1])
		os.Exit(1)
	}
}
