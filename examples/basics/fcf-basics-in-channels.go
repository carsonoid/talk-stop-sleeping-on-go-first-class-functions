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
	actions := make(chan func(string) string, 4)
	actions <- greet
	actions <- emphasize
	actions <- strings.ToUpper
	actions <- func(s string) string { return s + " - from a channel" }
	close(actions)

	name := "Tina"
	for action := range actions {
		name = action(name)
	}
	fmt.Println(name)
}

// END OMIT
