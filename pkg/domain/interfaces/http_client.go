package interfaces

// IHTTPClient is an interface for HTTP clients.
type IHTTPClient interface {
	Get(url string) (content string, err error)
}
