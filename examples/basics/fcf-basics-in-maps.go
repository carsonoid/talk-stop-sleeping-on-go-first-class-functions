package main

import (
	"fmt"
	"strings"
)

// START OMIT
func main() {
	actions := map[string]func(string) string{
		"greet":     greet,
		"emphasize": emphasize,
		"toUpper":   strings.ToUpper,
		"comment":   func(s string) string { return s + " - from a map" },
	}

	name := "Alice"
	name = actions["greet"](name)
	name = actions["emphasize"](name)
	name = actions["toUpper"](name)
	fmt.Println(name)
}

func greet(name string) string {
	return "Hello, " + name
}

func emphasize(s string) string {
	return s + "!"
}

// END OMIT
