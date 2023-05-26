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

const uri = "https://api.kanye.rest"

func quoteKanyeWest(censor bool) string {
	data := consumeAPI()
	var respObj Response
	json.Unmarshal(data, &respObj)
	quote := respObj.Quote

	if censor {
		return goaway.Censor(quote)
	}

	return quote
}

func consumeAPI() []byte {
	resp, err := http.Get(uri)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return data
}

func main() {
	fmt.Println(quoteKanyeWest(true))
}
