package ye_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	ye "github.com/imran-salim/ye/client"
)

func TestGetQuote(t *testing.T) {
	t.Run("GetQuote returns valid data", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"quote":"Buy property"}`))
		}))
		defer server.Close()

		ye.BaseUrl = server.URL
		t.Cleanup(func() { ye.BaseUrl = "https://api.kanye.rest" })

		data, err := ye.GetQuote()
		if err != nil {
			t.Fatalf("GetQuote() failed: %v", err)
		}

		if len(data) == 0 {
			t.Error("GetQuote() returned empty data")
		}
	})

	t.Run("GetQuote handles non-2xx response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusInternalServerError)
		}))
		defer server.Close()

		ye.BaseUrl = server.URL
		t.Cleanup(func() { ye.BaseUrl = "https://api.kanye.rest" })

		_, err := ye.GetQuote()
		if err == nil {
			t.Error("Expected error for 500 response, got nil")
		}
	})
}
