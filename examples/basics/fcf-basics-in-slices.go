package main

import (
	"fmt"
	"strings"
)

// START OMIT
func greet(name string) string {
	return "Hello, " + name
}

func emphasize(s string) string {
	return s + "!"
}

func main() {
	pipeline := []func(string) string{
		greet,
		emphasize,
		strings.ToUpper,
		func(s string) string { return s + " - from a slice" },
	}

	name := "Tami"
	for _, fn := range pipeline {
		name = fn(name)
	}
	fmt.Println(name)
}

// END OMIT
