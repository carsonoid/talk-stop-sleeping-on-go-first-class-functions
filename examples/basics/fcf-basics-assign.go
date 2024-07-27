package main

import "fmt"

// START OMIT
func main() {
	fn := printHello
	fn()

	fmt.Printf("--- fn ---\n")
	fmt.Printf("Type: %T\tValue: %v\n", fn, fn)

	fmt.Printf("--- printHello ---\n")
	fmt.Printf("Type: %T\tValue: %v\n", printHello, printHello)
}

func printHello() {
	fmt.Println("Hello, World!")
}

// END OMIT
