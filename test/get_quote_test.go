package main

import (
	"net/http"
	"os"
	"testing"
)

func testIsTheHttpRequestSuccessful(t *testing.T, url string) {
	resp, err := http.Get(url)
	if err != nil {
		t.Errorf("There was no response from the server: %s", url)
		os.Exit(1)
	}
	if resp.StatusCode > 299 {
		t.Errorf("Response failed with the status code: %d", resp.StatusCode)
	}
}

func TestGetQuote(t *testing.T) {
	url := "https://api.kanye.rest"
	testIsTheHttpRequestSuccessful(t, url)
}
