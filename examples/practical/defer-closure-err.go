package main

import (
	"errors"
	"fmt"
)

// START OMIT
func main() {
	err := possiblePanic()
	if err != nil {
		fmt.Println("error happened: ", err)
	} else {
		fmt.Println("no error")
	}
}

func possiblePanic() (err error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("!!! recovered from panic")
			err = errors.New("panic happened")
		}
	}()
	// panic("panic")
	return nil
}

// END OMIT
