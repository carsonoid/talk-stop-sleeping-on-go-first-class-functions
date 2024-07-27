package main

import "fmt"

// START OMIT
func main() {
	// assign and call
	fn := getFunc("hello")
	fn()

	// call returned function directly
	getFunc("goodbye")()
}

func printHello() {
	fmt.Println("Hello World!")
}

func printGoodbye() {
	fmt.Println("Goodbye World!")
}

func getFunc(s string) func() {
	if s == "hello" {
		return printHello
	}
	return printGoodbye
}

// END OMIT
