package ye

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

var BaseUrl = "https://api.kanye.rest"

func GetQuote() ([]byte, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(BaseUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("HTTP request failed: status %d", resp.StatusCode)
	}

	return body, nil
}
