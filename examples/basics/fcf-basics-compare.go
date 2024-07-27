package main

import "fmt"

// START OMIT
func printHello() {
	fmt.Println("Hello World!")
}

func main() {
	hello1 := printHello
	hello2 := printHello

	// 👍 valid!
	if hello1 == nil {
		fmt.Println("hello is nil")
	}

	// 🛑 INVALID! Functions cannot be compared
	//    even if they have the same signature
	if hello1 == hello2 {
		fmt.Println("hello1 and hello2 are equal")
	}
}

// END OMIT
