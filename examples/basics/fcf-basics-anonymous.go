package main

import "fmt"

// START OMIT
func main() {
	// named nested functions are not supported in Go :(
	// func printHello() {
	// 	fmt.Println("Hello World!")
	// }

	// but you can nest anonymous functions all you want!
	run := func() {
		sayHello := func() {
			fmt.Println("Hello World!")
		}
		sayHello()

		sayGoodbye := func() {
			fmt.Println("Goodbye World!")
		}
		sayGoodbye()
	}
	run()
	run()
}

// END OMIT
