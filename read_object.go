package main

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func ReadObject(repo Repo, sha SHA) (Object, error) {
	shaStr := sha.String()
	objectPath := path.Join(repo.GitDir, "objects", shaStr[:2], shaStr[2:])
	f, err := os.Open(objectPath)
	if err != nil {
		return Object{}, err
	}

	r, err := zlib.NewReader(f)
	if err != nil {
		return Object{}, err
	}

	fullData, err := io.ReadAll(r)
	if err != nil {
		return Object{}, err
	}

	parts := bytes.SplitN(fullData, []byte{'\x00'}, 2)
	assertEq(len(parts), 2, "object must be split to header and body only")
	header, body := string(parts[0]), parts[1]

	assertEq(len(parts), 2, "header must be split to kind and size only")
	headerParts := strings.Split(header, " ")
	kind, sizeStr := ObjectKind(headerParts[0]), headerParts[1]
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		return Object{}, fmt.Errorf("header/size part is not int: '%s', err: %w", sizeStr, err)
	}

	return Object{
		Kind: kind,
		Size: size,
		Body: body,
	}, nil
}
