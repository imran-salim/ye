package ye

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const url = "https://api.kanye.rest"

func GetQuote() []byte {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
