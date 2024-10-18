package main

import "fmt"

// START OMIT
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

func main() {
	// assign and call
	hello := getFunc("hello")
	hello()

	// call returned function directly
	getFunc("goodbye")()
}

// END OMIT
