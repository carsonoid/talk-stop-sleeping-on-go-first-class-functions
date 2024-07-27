package main

import (
	"fmt"
	"time"
)

// START ORIGINAL OMIT
func main() {
	fmt.Println(getNextHours(3))
}

func getNextHours(n int) []time.Time {
	now := time.Now()
	var s []time.Time
	for i := 0; i < n; i++ {
		s = append(s, now.Add(time.Hour*time.Duration(i+1)))
	}
	return s
}

// END ORIGINAL OMIT
