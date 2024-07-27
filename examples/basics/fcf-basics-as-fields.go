package main

// START OMIT
type client struct {
	name string
	do   func(string) // <- a function!
}

func (c *client) greet() {
	c.do(c.name)
}

func greet(name string) {
	println("Hello,", name)
}

func main() {
	clients := []client{
		{"Tami", func(name string) { println("Hello,", name+"! You rock!") }},
		{"Raul", greet},
	}

	for _, c := range clients {
		c.greet()
	}
}

// END OMIT
