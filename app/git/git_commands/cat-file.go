package gitcommands

import (
	"bytes"
	"compress/zlib"
	"io"
	"log"
	"os"
	"path/filepath"
)

//git cat-file -p hash
func Gitcatfile(flag string, commitSHA string) {

	dir,err := os.Getwd(); if err != nil {
		log.Fatal(err)
	}
	
	switch flag {
	case "-p" :
		folder := commitSHA[:2]
		file := commitSHA[2:]

		path := filepath.Join(dir , ".git" , "objects" , folder , file)
		
		if _,err := os.Stat(path); err != nil {
			log.Fatal(err)
		}

		//read the data 
		data,err := os.ReadFile(path); if err != nil {
			log.Fatal(err)
		}

		//decomress(inflate) data using zlib
		r,err := zlib.NewReader(io.NopCloser(bytes.NewReader(data))); if err != nil {
			log.Fatal(err)
		}


		defer r.Close()
		
		output,err := io.ReadAll(r); if err != nil {
			log.Fatal(err)
		}

		os.Stdout.Write(output)
		
	}
}
