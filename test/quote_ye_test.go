package main

import (
	"encoding/json"
	"strings"
	"testing"

	ye "github.com/imran-salim/ye/client"
)

func testIsThereAQuote(t *testing.T, respBody ye.Response) {
	if len(respBody.Quote) < 1 {
		t.Logf("The name/value pair in the body of the HTTP response does not contain a valid value")
		t.Fail()
	}
}

func testProfanityDetectorDictionaryEntries(t *testing.T, entries []string, typeOfEntries string) {
	if entries != nil {
		for index, element := range entries {
			if len(element) < 1 {
				t.Logf("There is an invalid element in the list of %s at index %d: \"%s\"", typeOfEntries, index, element)
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
			t.Logf("The quote contains profanity, but is not censored")
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
	testProfanityDetectorDictionaryEntries(t, []string{}, "profanities")
	testProfanityDetectorDictionaryEntries(t, []string{"button"}, "false positives")
	testProfanityDetectorDictionaryEntries(t, []string{"f#%k"}, "false negatives")
	testIsTheQuoteCensored(t, quote)
}
