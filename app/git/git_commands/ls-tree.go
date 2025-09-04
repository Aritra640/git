package gitcommands

import (
	"compress/zlib"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
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

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	folderPath := filepath.Join(dir, ".git", "objects", folder)
	if _, err = os.Stat(folderPath); err != nil {
		log.Println("folder not found")
		log.Fatal(err)
	}

	filePath := filepath.Join(folderPath, file)
	if _, err := os.Stat(filePath); err != nil {
		log.Println("file not found")
		log.Fatal(err)
	}

	fileContent, err := os.Open(filePath)
	if err != nil {
		log.Println("Error: cannot read file with hash ", g.sha)
		log.Fatal(err)
	}
	defer fileContent.Close()

	//create a zlib reader
	reader, err := zlib.NewReader(fileContent)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	decompressedData, err := io.ReadAll(reader)
	if err != nil {
		panic(err)
	}
	decompressedStr := string(decompressedData)
	parts := strings.Split(decompressedStr, "\x00")
	if len(parts) < 2 {
		log.Fatalf("expected at least 2 parts after splitting")
	}

	treeContent := parts[1:]
	var names []string
	for _, e := range treeContent {
		if strings.Contains(e, " ") {
			fields := strings.Split(e, " ")
			if len(fields) > 1 {
				names = append(names, fields[1])
			}
		}
	}
	for _, name := range names {
		fmt.Println(name)
	}

}
