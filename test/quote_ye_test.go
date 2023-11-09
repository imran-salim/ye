package main

import (
	"encoding/json"
	"strings"
	"testing"

	goaway "github.com/TwiN/go-away"
	ye "github.com/i8abyte/ye/client"
)

func testIsThereAQuote(t *testing.T, respBody ye.Response) {
	if len(respBody.Quote) < 1 {
		t.Errorf("The name/value pair in the body of the HTTP response does not contain a valid value")
	}
}

func testIsTheQuoteCensored(t *testing.T, quote string) {
	if goaway.IsProfane(quote) {
		censoredQuote := goaway.Censor(quote)

		if !strings.Contains(censoredQuote, "*") {
			t.Errorf("The quote contains profanity but is not censored")
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
