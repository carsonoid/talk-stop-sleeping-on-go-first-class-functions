package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// START ORIGINAL OMIT
func Test_getURL(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	t.Cleanup(s.Close)

	status, err := getURL(s.URL)
	if err != nil {
		t.Fatal(err)
	}

	if status != http.StatusOK {
		t.Errorf("got %d, want %d", status, http.StatusOK)
	}
}

// END ORIGINAL OMIT
