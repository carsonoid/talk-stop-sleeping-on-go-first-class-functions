package main

import "fmt"

// START OMIT
func main() {
	fn := printHello
	fmt.Printf("Type: %T\tValue: %v\n", fn, fn)

	say(fn, "World")
	say(printGoodbye, "World!")
}

func printHello() {
	fmt.Print("Hello ")
}

func printGoodbye() {
	fmt.Print("Goodbye ")
}

func say(fn func(), s string) {
	fn()
	fmt.Println(s)
}

// END OMIT
