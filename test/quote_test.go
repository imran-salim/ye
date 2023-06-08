package main

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	goaway "github.com/TwiN/go-away"
	ye "github.com/imran-salim/ye/api"
)

type Response struct {
	Quote string
}

func testRespObjType(t *testing.T) {
	data := ye.GetQuote()
	var respObj Response
	json.Unmarshal(data, &respObj)

	if reflect.TypeOf(respObj.Quote) != reflect.TypeOf(respObj.Quote) {
		t.Errorf("The data is not of the type %T", respObj.Quote)
	}
}

func testRespObjQuote(t *testing.T) {
	data := ye.GetQuote()
	var respObj Response
	json.Unmarshal(data, &respObj)

	if len(respObj.Quote) < 1 {
		t.Errorf("The quote in the response is empty: %s", respObj.Quote)
	}
}

func testCensor(t *testing.T) {
	data := ye.GetQuote()
	var respObj Response
	json.Unmarshal(data, &respObj)
	quote := respObj.Quote

	if goaway.IsProfane(quote) {
		censoredQuote := goaway.Censor(quote)

		if !strings.Contains(censoredQuote, "*") {
			t.Errorf("The quote could not be censored despite containing profanity")
		}
	}
}

func TestQuoteKanyeWest(t *testing.T) {
	testRespObjType(t)
	testRespObjQuote(t)
	testCensor(t)
}
