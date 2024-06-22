package main

import (
	"encoding/hex"
	"fmt"
)

const SHAEncodedLen = 40

type SHA []byte

func NewSHAFromString(str string) (SHA, error) {
	data, err := hex.DecodeString(str)
	if err != nil {
		return SHA{}, err
	}

	sha := SHA(data)
	sha.MustValidate()

	return sha, nil
}

func (s SHA) String() string {
	s.MustValidate()

	return hex.EncodeToString(s)
}

func (s SHA) MustValidate() {
	str := hex.EncodeToString(s)
	if len(str) != SHAEncodedLen {
		panic(fmt.Sprintf("incorrect sha length %d: %s", len(str), str))
	}
}
