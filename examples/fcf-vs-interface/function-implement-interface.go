package main

import (
	"fmt"
	"time"
)

// START OMIT
type greeter interface {
	SayHello(string)
}

type greeterFunc func(string)

func (g greeterFunc) SayHello(name string) {
	g(name)
}

func greet(g greeter, name string) {
	g.SayHello(name)
}

func main() {
	now := time.Now()
	greet(
		greeterFunc(func(name string) { fmt.Println("Hello,", name+"! It is", now) }),
		"Tami",
	)
}

// END OMIT
