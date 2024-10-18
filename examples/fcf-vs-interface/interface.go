package main

import (
	"fmt"
	"time"
)

// START OMIT
type greeter interface {
	SayHello(string)
}

type happyGreeter struct{ now time.Time }

func (g *happyGreeter) SayHello(name string) { fmt.Println("Hello,", name+"! It is", g.now) }

type sadGreeter struct{ now time.Time }

func (g *sadGreeter) SayHello(name string) { fmt.Println("Hi,", name+".... It is", g.now) }

func greet(g greeter, name string) {
	g.SayHello(name)
}

func main() {
	now := time.Now()
	greet(&happyGreeter{now}, "Tami")
	greet(&sadGreeter{now}, "Raul")
}

// END OMIT
