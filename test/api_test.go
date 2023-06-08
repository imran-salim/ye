package main

import (
	"reflect"
	"testing"

	ye "github.com/imran-salim/ye/api"
)

func TestGetQuoteEmptyData(t *testing.T) {
	data := ye.GetQuote()

	if len(data) < 1 {
		t.Errorf("The data has %d bytes", len(data))
	}
}

func TestGetQuoteDataType(t *testing.T) {
	var dataType []byte
	data := ye.GetQuote()

	if reflect.TypeOf(data) != reflect.TypeOf(dataType) {
		t.Errorf("The data is not of the type %T", dataType)
	}
}
