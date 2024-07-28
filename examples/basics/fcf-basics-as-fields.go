package main

// START OMIT
func main() {
	clients := []client{
		{"Alice", func(name string) { println("Hello,", name+"! You rock!") }},
		{"Bob", greet},
	}

	for _, c := range clients {
		c.greet()
	}
}

type client struct {
	name string
	do   func(string)
}

func (c *client) greet() {
	c.do(c.name)
}

func greet(name string) {
	println("Hello,", name)
}

// END OMIT
