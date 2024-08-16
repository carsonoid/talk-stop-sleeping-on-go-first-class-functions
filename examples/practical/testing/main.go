package main

import (
	"io"
	"net/http"
	"time"
)

func main() {
	url := "http://www.google.com"
	status, err := getURL(url)
	if err != nil {
		panic(err)
	}
	println(status)
}

// START ORIGINAL OMIT
func getURL(url string) (int, error) {
	c := &http.Client{Timeout: 5 * time.Second}
	resp, err := c.Get(url)
	if err != nil {
		return 0, err
	}
	io.Copy(io.Discard, resp.Body)
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

// END ORIGINAL OMIT
