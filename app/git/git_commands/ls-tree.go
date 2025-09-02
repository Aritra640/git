package gitcommands

import (
	"log"
	"os"
	"path/filepath"
)

type Git_ls_tree struct {
	flag string
	sha  string
}

func (g *Git_ls_tree) Init() {
	if len(os.Args) == 4 {
		g.flag = os.Args[2]
		g.sha = os.Args[3]

	} else {
		g.flag = ""
		g.sha = os.Args[2]
	}
}


func (g *Git_ls_tree) Execute() {

	folder := g.sha[:2]
	file := g.sha[2:]

	dir,err := os.Getwd(); if err != nil {
		log.Fatal(err)
	}

	folderPath := filepath.Join(dir, ".git", "objects", folder)
	if _,err = os.Stat(folderPath); err != nil {
		log.Println("folder not found")
		log.Fatal(err)
	}

	filePath := filepath.Join(folderPath , file) 
	if _,err := os.Stat(filePath); err != nil {
		log.Println("file not found")
		log.Fatal(err)
	}
}
