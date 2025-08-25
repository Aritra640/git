package gitcommands

import (
	"log"
	"os"
	"path/filepath"
)

// Create git directory
func Gitinit() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	path := filepath.Join(dir, ".git")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal(err)
	}

	path = filepath.Join(dir, ".git", "objects")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal(err)
	}

	path = filepath.Join(dir, ".git", "refs")
	err = os.MkdirAll(path, 0755)
	if err != nil {
		log.Fatal(err)
	}

	//Write HEAD file
	headPath := filepath.Join(dir, ".git", "HEAD")
	headContent := []byte("ref: refs/heads/main\n")
	err = os.WriteFile(headPath , headContent , 0644); if err != nil {
		log.Fatal(err)
	}

	log.Println("Intialized git directory")
}
