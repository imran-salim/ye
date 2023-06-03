package main

import (
	"encoding/json"

	goaway "github.com/TwiN/go-away"
)

type Response struct {
	Quote string
}

func quoteKanyeWest(censored bool) string {
	data := consumeAPI()
	var respObj Response
	json.Unmarshal(data, &respObj)
	quote := respObj.Quote

	if censored {
		return goaway.Censor(quote)
	}

	return quote
}
