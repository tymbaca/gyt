package main

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
)

func HashObject(objectKind ObjectKind, r io.Reader) (SHA, error) {
	fileData, err := io.ReadAll(r)
	if err != nil {
		return SHA{}, err
	}

	// <kind> + ' ' + len(file) + '\x00' + file
	builder := bytes.Buffer{}
	_, err = builder.WriteString(string(objectKind))
	if err != nil {
		return SHA{}, err
	}
	err = builder.WriteByte(' ')
	if err != nil {
		return SHA{}, err
	}
	_, err = builder.WriteString(fmt.Sprint(len(fileData)))
	if err != nil {
		return SHA{}, err
	}
	err = builder.WriteByte('\x00')
	if err != nil {
		return SHA{}, err
	}
	_, err = builder.Write(fileData)
	if err != nil {
		return SHA{}, err
	}

	data := builder.Bytes()
	// Logf("%s\n", data)

	sha := sha1.Sum(data)
	return sha[:], nil
}
