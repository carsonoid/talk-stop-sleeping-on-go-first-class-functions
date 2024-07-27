package main

import "fmt"

// START OMIT
func printHello() {
	fmt.Println("Hello, World!")
}

func main() {
	hello := printHello
	hello()

	fmt.Printf("--- fn ---\n")
	fmt.Printf("Type: %T\tValue: %v\n", hello, hello)

	fmt.Printf("--- printHello ---\n")
	fmt.Printf("Type: %T\tValue: %v\n", printHello, printHello)
}

// END OMIT
