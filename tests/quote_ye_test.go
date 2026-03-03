package ye

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	ye "github.com/imran-salim/ye/client"
)

func TestQuoteYe(t *testing.T) {
	t.Run("Returns non-empty quote", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"quote":"Buy property"}`))
		}))
		defer server.Close()

		ye.BaseUrl = server.URL
		t.Cleanup(func() { ye.BaseUrl = "https://api.kanye.rest" })

		quote := ye.QuoteYe()
		if quote == "" {
			t.Error("QuoteYe() returned empty string")
		}
	})

	t.Run("Quote is properly formatted", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`{"quote":"Buy property"}`))
		}))
		defer server.Close()

		ye.BaseUrl = server.URL
		t.Cleanup(func() { ye.BaseUrl = "https://api.kanye.rest" })

		quote := ye.QuoteYe()
		if !strings.HasPrefix(quote, "\"") || !strings.HasSuffix(quote, "\"") {
			t.Errorf("Quote is not properly quoted: %q", quote)
		}
	})

	t.Run("profanity detector handles custom words", func(t *testing.T) {
		detector := ye.ProfanityDetector([]string{"test"}, []string{"button"}, []string{"f#%k"})

		if !detector.IsProfane("test word") {
			t.Error("Custom profanity not detected")
		}

		if detector.IsProfane("button") {
			t.Error("False positive not handled")
		}
	})
}
