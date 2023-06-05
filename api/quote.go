package api

import (
	"encoding/json"

	goaway "github.com/TwiN/go-away"
)

type Response struct {
	Quote string
}

func QuoteKanyeWest(censored bool) string {
	data := ConsumeAPI()
	var respObj Response
	json.Unmarshal(data, &respObj)
	quote := respObj.Quote

	if censored {
		return goaway.Censor(quote)
	}

	return quote
}
