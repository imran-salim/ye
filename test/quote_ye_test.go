package ye

import (
	"strings"
	"testing"

	ye "github.com/imran-salim/ye/client"
)

func TestQuoteYe(t *testing.T) {
	t.Run("returns non-empty quote", func(t *testing.T) {
		quote := ye.QuoteYe()
		if quote == "" {
			t.Error("QuoteYe() returned empty string")
		}
	})

	t.Run("quote is properly formatted", func(t *testing.T) {
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
