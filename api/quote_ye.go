package ye

import (
	"encoding/json"

	goaway "github.com/TwiN/go-away"
)

type Response struct {
	Quote string
}

func QuoteKanyeWest(censored bool) string {
	data := GetQuote()
	var respObj Response
	json.Unmarshal(data, &respObj)
	quote := respObj.Quote

	if goaway.IsProfane(quote) && censored {
		return goaway.Censor(quote)
	}

	return quote
}
