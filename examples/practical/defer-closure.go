package main

import (
	"fmt"
)

// START OMIT
func main() {
	s := possiblePanic()
	fmt.Printf("s: %q\n", s)
}

func possiblePanic() (s string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("!!! recovered from panic")
			s = "default"
		}
	}()
	panic("panic")
}

// END OMIT
