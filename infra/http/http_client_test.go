package http

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewHTTPClient(t *testing.T) {
	timeout := 2 * time.Second
	client := NewHTTPClient(timeout)

	assert.NotNil(t, client)
	assert.NotNil(t, client.client)
	assert.Equal(t, timeout, client.client.Timeout)
}

func TestHTTPClient_Get_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello"))
	}))
	defer server.Close()

	client := NewHTTPClient(2 * time.Second)
	content, err := client.Get(server.URL)

	assert.NoError(t, err)
	assert.Equal(t, "Hello", content)
}

func TestHTTPClient_Get_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := NewHTTPClient(2 * time.Second)
	content, err := client.Get(server.URL)

	assert.Error(t, err)
	assert.Empty(t, content)
}

func TestHTTPClient_Get_Timeout(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(1100 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Delayed response"))
	}))
	defer server.Close()

	client := NewHTTPClient(1 * time.Second)
	content, err := client.Get(server.URL)

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Timeout")
	assert.Empty(t, content)
}
