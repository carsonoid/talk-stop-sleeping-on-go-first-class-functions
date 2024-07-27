package main

import (
	"fmt"
	"time"
)

// START OMIT
func main() {
	triggerUpdateChan := getTriggerUpdateChan()
	defer close(triggerUpdateChan)

	go processUpdateTriggers(triggerUpdateChan)

	// simulate a trigger event (file change, network event, etc.)
	triggerUpdateChan <- struct{}{}
	fmt.Println("Triggered Update 1")

	triggerUpdateChan <- struct{}{}
	fmt.Println("Triggered Update 2")

	triggerUpdateChan <- struct{}{}
	fmt.Println("Triggered Update 3")

	time.Sleep(2 * time.Second)
}

func getTriggerUpdateChan() chan struct{} {
	queue := make(chan struct{}, 1)
	return queue
}

func processUpdateTriggers(queue <-chan struct{}) {
	for range queue {
		handleUpdate()
	}
}

func handleUpdate() {
	fmt.Println("Handling Update...")
	time.Sleep(1 * time.Second)
}

// END OMIT
