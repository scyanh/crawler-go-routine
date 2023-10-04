package entities

// WorkerPool represents the interface for a worker pool handling webpage crawling.
type WorkerPool interface {
	Start()
	AddJob(webPage *WebPage)
}
