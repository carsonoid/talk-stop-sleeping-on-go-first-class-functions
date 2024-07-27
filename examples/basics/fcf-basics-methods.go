package main

import "fmt"

// START OMIT
func useGreeter(fn func(string), s string) {
	fn(s)
}

type greeter struct {
	msg string
}

func (g greeter) greet(s string) {
	fmt.Println(g.msg, s)
}

func main() {
	g := greeter{msg: "Hello "}
	useGreeter(g.greet, "World")
	greet := g.greet
	useGreeter(greet, "World")
}

// END OMIT
