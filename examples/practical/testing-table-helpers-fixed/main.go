package main

import (
	"context"
)

// START ORIGINAL OMIT
func main() {
	hello, err := getHello(context.Background())
	if err != nil {
		panic(err)
	}
	println(hello)
}

func getHello(ctx context.Context) (string, error) {
	if err := ctx.Err(); err != nil {
		return "", nil
	}

	return "Hello, world!", nil
}

// END ORIGINAL OMIT
