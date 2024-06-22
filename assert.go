package main

import "fmt"

func assert(val bool, msg string) {
	if !val {
		panic(msg)
	}
}

func assertEq[T comparable](got, want T, msg string) {
	if got != want {
		panic(fmt.Sprintf("%s, got: %v, want: %v", msg, got, want))
	}
}
