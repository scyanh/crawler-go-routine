package interfaces

// IHTTPClient is an interface for HTTP client.
type IHTTPClient interface {
	Get(url string) (content string, err error)
}
