package main

import (
	"fmt"
	"strings"
)

// START OMIT
func main() {
	actions := make(chan func(string) string, 4)
	actions <- greet
	actions <- emphasize
	actions <- strings.ToUpper
	actions <- func(s string) string { return s + " - from a channel" }
	close(actions)

	name := "Alice"
	for action := range actions {
		name = action(name)
	}
	fmt.Println(name)
}

func greet(name string) string {
	return "Hello, " + name
}

func emphasize(s string) string {
	return s + "!"
}

// END OMIT
