package ye

import (
	"encoding/json"

	goaway "github.com/TwiN/go-away"
)

type Response struct {
	Quote string
}

func QuoteYe(censored bool) string {
	data := GetQuote()
	var respBody Response
	json.Unmarshal(data, &respBody)
	quote := "\"" + respBody.Quote + "\""
	profanityDetector := CustomProfanityDetector()

	if profanityDetector.IsProfane(quote) && censored {
		return profanityDetector.Censor(quote)
	}

	return quote
}

func CustomProfanityDetector() *goaway.ProfanityDetector {
	profanityDetector := goaway.NewProfanityDetector()
	customFalsePositives := goaway.DefaultFalsePositives
	customFalsePositives = append(customFalsePositives, "therapeutic")
	profanityDetector.WithCustomDictionary(goaway.DefaultProfanities, customFalsePositives, goaway.DefaultFalseNegatives)
	return profanityDetector
}
