package main

import (
	"encoding/json"
	"testing"

	ye "github.com/i8abyte/ye/api"
)

func testIsRespBodyEmpty(t *testing.T, respBody ye.Response) {
	if len(respBody.Quote) < 1 {
		t.Errorf("The name/value pair in the body of the HTTP response does not contain a value")
	}
}

func TestGetQuote(t *testing.T) {
	data := ye.GetQuote()
	var respBody ye.Response
	json.Unmarshal(data, &respBody)

	testIsRespBodyEmpty(t, respBody)
}
