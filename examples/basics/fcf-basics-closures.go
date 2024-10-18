package main

import (
	"fmt"
	"time"
)

// START OMIT
func getHello(now time.Time) func() {
	return func() {
		fmt.Println("Hello, World! It's", now)
	}
}

func callTwice(fn func()) {
	fn()
	fn()
}

func main() {
	hello := getHello(time.Now())
	hello()
	callTwice(hello)
	callTwice(func() {
		fmt.Print("closure a closure: ")
		hello()
	})
}

// END OMIT
