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
	return getNextHoursFn(time.Now, n)
}

func getNextHoursFn(getNow func() time.Time, n int) []time.Time {
	now := getNow()
	var s []time.Time
	for i := 0; i < n; i++ {
		s = append(s, now.Add(time.Hour*time.Duration(i+1)))
	}
	return s
}

// END ORIGINAL OMIT
