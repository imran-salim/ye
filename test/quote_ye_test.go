package main

import (
	"encoding/json"
	"strings"
	"testing"

	ye "github.com/i8abyte/ye/client"
)

func testIsThereAQuote(t *testing.T, respBody ye.Response) {
	if len(respBody.Quote) < 1 {
		t.Fatalf("The name/value pair in the body of the HTTP response does not contain a valid value")
	}
}

func testIsTheQuoteCensored(t *testing.T, quote string) {
	profanityDetector := ye.CustomProfanityDetector()

	if profanityDetector.IsProfane(quote) {
		censoredQuote := profanityDetector.Censor(quote)

		if !strings.Contains(censoredQuote, "*") {
			t.Fatalf("The quote contains profanity but is not censored")
		}
	}
}

func TestQuoteYe(t *testing.T) {
	data := ye.GetQuote()
	var respBody ye.Response
	json.Unmarshal(data, &respBody)
	testIsThereAQuote(t, respBody)
	quote := "\"" + respBody.Quote + "\""
	testIsTheQuoteCensored(t, quote)
}
