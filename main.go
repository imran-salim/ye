package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	goaway "github.com/TwiN/go-away"
)

type Response struct {
	Quote string
}

func quoteKanyeWest(censor bool) string {
	uri := "https://api.kanye.rest"

	res, err := http.Get(uri)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(data, &responseObject)

	quote := responseObject.Quote

	if censor {
		return goaway.Censor(quote)
	}

	return quote
}

func main() {
	fmt.Println(quoteKanyeWest(true))
}
