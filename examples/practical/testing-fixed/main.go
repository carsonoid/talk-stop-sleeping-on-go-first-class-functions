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

	c := &http.Client{Timeout: 5 * time.Second}
	status, err = getURLWithClient(c, url)
	if err != nil {
		panic(err)
	}
	println(status)

	status, err = getURLWithGlobal(url)
	if err != nil {
		panic(err)
	}
	println(status)

	status, err = getURLWithFunc(c.Get, url)
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

// START WITH_CLIENT OMIT
func getURLWithClient(c *http.Client, url string) (int, error) {
	resp, err := c.Get(url)
	if err != nil {
		return 0, err
	}
	io.Copy(io.Discard, resp.Body)
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

// END WITH_CLIENT OMIT

// START WITH_INTERFACE OMIT
type getter interface {
	Get(url string) (resp *http.Response, err error)
}

func getURLWithInterface(g getter, url string) (int, error) {
	resp, err := g.Get(url)
	if err != nil {
		return 0, err
	}
	io.Copy(io.Discard, resp.Body)
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

// END WITH_INTERFACE OMIT

// START WITH_GLOBAL OMIT
var globalGetFunc = (&http.Client{Timeout: 5 * time.Second}).Get

func getURLWithGlobal(url string) (int, error) {
	resp, err := globalGetFunc(url)
	if err != nil {
		return 0, err
	}
	io.Copy(io.Discard, resp.Body)
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

// END WITH_GLOBAL OMIT

// START WITH_FUNC OMIT
type getFunc func(url string) (resp *http.Response, err error)

func getURLWithFunc(get getFunc, url string) (int, error) {
	resp, err := get(url)
	if err != nil {
		return 0, err
	}
	io.Copy(io.Discard, resp.Body)
	defer resp.Body.Close()
	return resp.StatusCode, nil
}

// END WITH_FUNC OMIT
