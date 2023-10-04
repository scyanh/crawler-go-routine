package http

import (
	"io/ioutil"
	"net/http"
	"time"
)

type HTTPClient struct {
	client *http.Client
}

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		client: &http.Client{
			Timeout: time.Second * 10, // Definir un tiempo de espera de 10 segundos para las solicitudes
		},
	}
}

func (c *HTTPClient) Get(url string) (content string, err error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// Asegúrate de que recibimos un código de estado 200 OK
	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}
