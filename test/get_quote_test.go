package ye

import (
	"encoding/json"
	"net/http"
	"testing"

	ye "github.com/imran-salim/ye/client"
)

func TestGetQuote(t *testing.T) {
	t.Run("API endpoint is reachable", func(t *testing.T) {
		url := "https://api.kanye.rest"
		resp, err := http.Get(url)
		if err != nil {
			t.Fatalf("There was no response from the server: %s", url)
		}
		defer resp.Body.Close()

		if resp.StatusCode > 299 {
			t.Fatalf("Response failed with the status code: %d", resp.StatusCode)
		}
	})

	t.Run("GetQuote returns valid data", func(t *testing.T) {
		data, err := ye.GetQuote()
		if err != nil {
			t.Fatalf("GetQuote() failed: %v", err)
		}

		if len(data) == 0 {
			t.Error("GetQuote() returned empty data")
		}
	})

	t.Run("GetQuote returns valid JSON", func(t *testing.T) {
		data, err := ye.GetQuote()
		if err != nil {
			t.Fatalf("GetQuote() failed: %v", err)
		}

		var respBody ye.Response
		err = json.Unmarshal(data, &respBody)
		if err != nil {
			t.Fatalf("Failed to unmarshal JSON response: %v", err)
		}

		if respBody.Quote == "" {
			t.Error("Resposne contains empty quote")
		}
	})
}
