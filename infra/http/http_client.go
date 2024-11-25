package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type HTTPClient struct {
	client *http.Client
}

// NewHTTPClient creates a new HTTPClient with the given timeout.
func NewHTTPClient(timeout time.Duration) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: timeout,
		},
	}
}

// Get performs a GET request to the given URL and returns the response body.
func (c *HTTPClient) Get(url string) (content string, err error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
