package main

import "fmt"

// START OMIT
func main() {
	printHello := func() {
		fmt.Println("Hello World!")
	}
	printHello()

	printGoodbye := func() {
		fmt.Println("Goodbye World!")
	}
	printGoodbye()
}

// END OMIT
