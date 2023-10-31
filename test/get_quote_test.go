package main

import (
	"net/http"
	"os"
	"testing"
)

func testIsHttpRequestSuccessful(t *testing.T, uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		t.Errorf("There was no response from the server: %s", uri)
		os.Exit(1)
	}
	if resp.StatusCode > 299 {
		t.Errorf("Response failed with the status code: %d", resp.StatusCode)
	}
}

func TestGetQuote(t *testing.T) {
	uri := "https://api.kanye.rest"
	testIsHttpRequestSuccessful(t, uri)
}
