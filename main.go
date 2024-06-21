package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	args := os.Args
	fmt.Println(args)

	if len(args) < 3 {
		log.Fatalf(`type some commands: "hash-object" or "cat-file" and target file`)
	}

	switch args[1] {
	case "cat-file":
	case "hash-object":
		f, err := os.Open(args[2])
		if err != nil {
			panic(err)
		}

		sha, err := HashObject(f)
		if err != nil {
			panic(err)
		}

		fmt.Println(sha)
	}
}

type SHA [sha1.Size]byte

func (s SHA) String() string {
	return hex.EncodeToString(s[:])
}

func HashObject(r io.Reader) (SHA, error) {
	data, err := io.ReadAll(r)
	fmt.Printf("%s\n", data)
	if err != nil {
		return SHA{}, err
	}

	return sha1.Sum(data), nil
}
