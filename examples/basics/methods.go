package main

import (
	"fmt"
)

func main() {
	g := greeter{"Hello,"}
	g.sayHello("World")
}

type greeter struct {
	msg string
}

func (g *greeter) sayHello(user string) {
	fmt.Println(g.msg, user+"!")
}
