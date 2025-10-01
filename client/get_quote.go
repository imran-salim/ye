package ye

import (
	"fmt"
	"io"
	"net/http"
)

const baseUrl = "https://api.kanye.rest"

func GetQuote() ([]byte, error) {
	resp, err := http.Get(baseUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if resp.StatusCode > 299 {
		return nil, fmt.Errorf("HTTP request failed: status %d", resp.StatusCode)
	}
	if err != nil {
		return nil, err
	}

	return body, nil
}
