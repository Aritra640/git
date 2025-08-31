package gitcommands

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

type GitHashObject struct {
	flag string
	file string
}

func (h *GitHashObject) Init() {

	if len(os.Args) == 4 {
		h.flag = os.Args[2]
		h.file = os.Args[3]

	} else {
		h.file = os.Args[2]
		h.flag = ""
	}
}

func (h *GitHashObject) Execute() {

	path, err := filepath.Abs(h.file)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := os.Stat(path); err != nil {
		if err == os.ErrNotExist {
			os.Stdout.Write([]byte(fmt.Sprintf("fatal: could not open '%v' for reading: No such file or directory", path)))
		} else {

			log.Fatal(err)
		}
	}

	filecontents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	filelength := len(filecontents)

	blob := fmt.Sprintf("blob %v\x00%v", filelength, string(filecontents))

	hash := sha1.New()
	hash.Write([]byte(blob))
	hashbytes := hash.Sum(nil)
	hexHash := hex.EncodeToString(hashbytes)

	if h.flag == "-w" {
		folder := hexHash[:2]
		newfile := hexHash[2:]

		dir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		cpath := filepath.Join(dir, ".git", "objects", folder)
		err = os.MkdirAll(cpath, 0755)
		if err != nil {
			log.Fatal(err)
		}

		var buf bytes.Buffer
		zw := zlib.NewWriter(&buf)
		_, err = zw.Write([]byte(blob))
		if err != nil {
			log.Fatal(err)
		}

		zw.Close()
		compressedData := buf.Bytes()

		outputPath := filepath.Join(cpath, newfile)
		if err = os.WriteFile(outputPath, compressedData, 0644); err != nil {
			log.Fatal(err)
		}
	}

	os.Stdout.Write([]byte(hexHash))
}
