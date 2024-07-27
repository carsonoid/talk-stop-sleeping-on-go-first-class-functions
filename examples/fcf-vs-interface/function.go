package main

import (
	"fmt"
	"time"
)

// START OMIT
func greet(fn func(string), name string) {
	fn(name)
}

func main() {
	now := time.Now()
	greet(
		func(name string) { fmt.Println("Hello,", name+"!", "It is", now) },
		"Tami",
	)
	greet(
		func(name string) { fmt.Println("Hi,", name+"....", "It is", now) },
		"Raul",
	)
}

// END OMIT
