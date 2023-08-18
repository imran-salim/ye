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
	var respBody Response
	json.Unmarshal(data, &respBody)
	quote := respBody.Quote

	if goaway.IsProfane(quote) && censored {
		return "\"" + goaway.Censor(quote) + "\""
	}

	return "\"" + quote + "\""
}
