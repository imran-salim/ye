package ye

import (
	"io"
	"log"
	"net/http"
)

const uri = "https://api.kanye.rest"

func GetQuote() []byte {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
	}

	return body
}
