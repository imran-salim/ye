package ye

import (
	"io"
	"log"
	"net/http"
)

const uri = "https://api.kanye.rest"

func GetQuote() []byte {
	res, err := http.Get(uri)

	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()

	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}

	if err != nil {
		log.Fatal(err)
	}

	return body
}
