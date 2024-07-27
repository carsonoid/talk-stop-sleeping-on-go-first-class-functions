package main

import "fmt"

// START OMIT
func main() {
	printHello := func() {
		fmt.Println("Hello World!")
	}
	printHello()

	// functions can also be called without
	// assigning them to a variable
	func() {
		fmt.Println("Goodbye World!")
	}()
}

// END OMIT
