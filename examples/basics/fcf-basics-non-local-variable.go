package main

import "fmt"

// START OMIT
var bang = "!"

func main() {
	world := "World"

	run := func() {
		hello := "Hello"

		sayHello := func() {
			fmt.Println(hello, world+bang)
		}
		sayHello()
	}
	run()
}

// END OMIT
