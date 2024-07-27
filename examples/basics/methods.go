package main

import (
	"fmt"
)

func main() {
	g := greeter{}
	g.sayHello("World")
}

type greeter struct {
}

func (g *greeter) sayHello(user string) {
	fmt.Println("Hello, ", user, "!")
}
