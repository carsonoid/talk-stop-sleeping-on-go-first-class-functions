package main

import "fmt"

// START OMIT
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

func main() {
	hello := printHello

	say(hello, "World")
	say(printGoodbye, "World!")
}

// END OMIT
