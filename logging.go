package main

import "fmt"

const Logging = true

func Log(args ...any) {
	if Logging {
		fmt.Println(args...)
	}
}

func Logf(format string, args ...any) {
	if Logging {
		fmt.Printf(format, args...)
	}
}
