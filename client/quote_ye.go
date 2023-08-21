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

	var resBody Response
	json.Unmarshal(data, &resBody)

	quote := "\"" + resBody.Quote + "\""

	if goaway.IsProfane(quote) && censored {
		return goaway.Censor(quote)
	}

	return quote
}
