package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	fn := getHello(time.Now())
	fn()
	time.Sleep(1 * time.Second)
	fn()
}

func getHello(now time.Time) func() {
	return func() {
		fmt.Println("Hello, World! It's", now)
	}
}

// END OMIT
