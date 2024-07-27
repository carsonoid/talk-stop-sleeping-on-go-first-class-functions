package main

import "fmt"

// START OMIT
func main() {
	tryQueue, cleanup := getQueue()
	defer cleanup()
	tryQueue("Hello, World!")
	tryQueue("Hello, World!")
}

// START GETQUEUE OMIT
func getQueue() (tryQueue func(string), cleanup func())
	queue := make(chan string, 1)
	return func(msg string) {
		select {
		case queue <- msg:
		default:
			fmt.Println("Queue is full, dropping message:", msg)
		}
	}, func() {
		close(queue)
	}
}
// END GETQUEUE OMIT

func processQueue(queue chan string) {
	for msg, ok := <-queue; ok; {
		fmt.Println(msg)
	}
}

// END OMIT
