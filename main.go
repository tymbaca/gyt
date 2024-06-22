package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args
	// Log(args)

	repo, err := OpenRepo()
	if err != nil {
		panic(err)
	}

	if len(args) < 3 {
		log.Fatalf(`type some commands: "hash-object" or "cat-file" and target file`)
	}

	switch args[1] {
	case "cat-file":
		mode := args[2]
		shaStr := args[3]
		CatFile(repo, mode, shaStr)

	case "hash-object":
		f := MustOpen(args[2])
		sha, err := HashObject(Blob, f)
		if err != nil {
			panic(err)
		}

		fmt.Println(sha.String())
	case "read-object":
		sha, err := NewSHAFromString(args[2])
		if err != nil {
			panic(err)
		}

		o, err := ReadObject(repo, sha)
		if err != nil {
			panic(err)
		}

		Logf("%s", o.Body)
	}
}

func MustOpen(path string) *os.File {
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	return f
}
