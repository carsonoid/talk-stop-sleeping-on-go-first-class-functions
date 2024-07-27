package main

import (
	"fmt"
	"strings"
)

// START OMIT
type stringFunc func(string) string

func callAndPrint(fn stringFunc, s string) {
	fmt.Println(fn(s))
}

func main() {
	var greet stringFunc = func(s string) string {
		return "Hello, " + s
	}

	emphasize := func(s string) string {
		return s + "!"
	}

	callAndPrint(greet, "Tami")
	callAndPrint(emphasize, "Raul")
	callAndPrint(strings.ToUpper, "Joster")
}

// END OMIT
