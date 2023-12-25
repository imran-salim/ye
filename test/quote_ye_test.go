package main

import (
	"encoding/json"
	"strings"
	"testing"

	ye "github.com/i8abyte/ye/client"
)

func testIsThereAQuote(t *testing.T, respBody ye.Response) {
	if len(respBody.Quote) < 1 {
		t.Logf("The name/value pair in the body of the HTTP response does not contain a valid value")
		t.Fail()
	}
}

func testProfanityDetecorDictionaryProfanities(t *testing.T, profanities []string) {
	if profanities != nil {
		for index, element := range profanities {
			if len(element) < 1 {
				t.Logf("There is an invalid element in the list of profanities at index %d: \"%s\"", index, element)
				t.Fail()
			}
		}
	}
}

func testProfanityDetecorDictionaryFalsePositives(t *testing.T, falsePositives []string) {
	if falsePositives != nil {
		for index, element := range falsePositives {
			if len(element) < 1 {
				t.Logf("There is an invalid element in the list of false positives at index %d: \"%s\"", index, element)
				t.Fail()
			}
		}
	}
}

func testProfanityDetecorDictionaryFalseNegatives(t *testing.T, falseNegatives []string) {
	if falseNegatives != nil {
		for index, element := range falseNegatives {
			if len(element) < 1 {
				t.Logf("There is an invalid element in the list of false negatives at index %d: \"%s\"", index, element)
				t.Fail()
			}
		}
	}
}

func testIsTheQuoteCensored(t *testing.T, quote string) {
	profanityDetector := ye.ProfanityDetector([]string{}, []string{"button"}, []string{"f#%k"})

	if profanityDetector.IsProfane(quote) {
		censoredQuote := profanityDetector.Censor(quote)

		if !strings.Contains(censoredQuote, "*") {
			t.Logf("The quote contains profanity but is not censored")
			t.Fail()
		}
	}
}

func TestQuoteYe(t *testing.T) {
	data := ye.GetQuote()
	var respBody ye.Response
	json.Unmarshal(data, &respBody)
	quote := "\"" + respBody.Quote + "\""

	testIsThereAQuote(t, respBody)
	testProfanityDetecorDictionaryProfanities(t, []string{})
	testProfanityDetecorDictionaryFalsePositives(t, []string{"button"})
	testProfanityDetecorDictionaryFalseNegatives(t, []string{"f#%k"})
	testIsTheQuoteCensored(t, quote)
}
