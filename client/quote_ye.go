package ye

import (
	"encoding/json"

	goaway "github.com/TwiN/go-away"
)

type Response struct {
	Quote string `json:"quote"`
}

func QuoteYe() string {
	data, err := GetQuote()
	if err != nil {
		return "Error fetching quote: " + err.Error()
	}

	var respBody Response
	err = json.Unmarshal(data, &respBody)
	if err != nil {
		return "Error parsing quote data"
	}

	quote := "\"" + respBody.Quote + "\""

	profanityDetector := ProfanityDetector([]string{}, []string{"button"}, []string{"f#%k"})
	if profanityDetector.IsProfane(quote) {
		return profanityDetector.Censor(quote)
	}

	return quote
}

func ProfanityDetector(customProfanities []string, customFalsePositives []string, customFalseNegatives []string) *goaway.ProfanityDetector {
	profanityDetector := goaway.NewProfanityDetector().WithSanitizeLeetSpeak(false).WithSanitizeSpecialCharacters(false)
	profanities, falsePositives, falseNegatives := goaway.DefaultProfanities, goaway.DefaultFalsePositives, goaway.DefaultFalseNegatives

	if customProfanities != nil {
		profanities = append(profanities, customProfanities...)
	}
	if customFalsePositives != nil {
		falsePositives = append(falsePositives, customFalsePositives...)
	}
	if customFalseNegatives != nil {
		falseNegatives = append(falseNegatives, customFalseNegatives...)
	}

	return profanityDetector.WithCustomDictionary(profanities, falsePositives, falseNegatives)
}
