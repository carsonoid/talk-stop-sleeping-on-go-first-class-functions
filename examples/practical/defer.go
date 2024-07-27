package main

import (
	"fmt"
)

// START OMIT
func main() {
	s := possiblePanic()
	fmt.Printf("s: %q\n", s)
}

func possiblePanic() string {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("!!! recovered from panic")
		}
	}()
	panic("panic")
	return "unreachable"
}

// END OMIT
