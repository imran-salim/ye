package ye

import (
	"encoding/json"

	goaway "github.com/TwiN/go-away"
)

type Response struct {
	Quote string `json:"quote"`
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
	profanityDetector := goaway.NewProfanityDetector().WithSanitizeLeetSpeak(false).WithSanitizeSpecialCharacters(false)

	customFalsePositives := goaway.DefaultFalsePositives
	customFalsePositives = append(customFalsePositives, "button")

	customFalseNegatives := goaway.DefaultFalseNegatives
	customFalseNegatives = append(customFalseNegatives, "f#%k")

	profanityDetector.WithCustomDictionary(goaway.DefaultProfanities, customFalsePositives, customFalseNegatives)

	return profanityDetector
}
