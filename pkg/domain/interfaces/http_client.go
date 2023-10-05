package interfaces

type IHTTPClient interface {
	Get(url string) (content string, err error)
}
