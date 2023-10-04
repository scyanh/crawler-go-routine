package worker

import (
	"fmt"
	"github.com/scyanh/crawler/pkg/domain"
	"github.com/scyanh/crawler/pkg/domain/entities"
	"github.com/scyanh/crawler/pkg/infra/http"
)

type GoroutineWorkerPool struct {
	workers int
	client  *http.HTTPClient
	repo    domain.Repository
	jobs    chan *entities.WebPage
}

func NewGoroutineWorkerPool(n int, client *http.HTTPClient, repo domain.Repository) *GoroutineWorkerPool {
	return &GoroutineWorkerPool{
		workers: n,
		client:  client,
		repo:    repo,
		jobs:    make(chan *entities.WebPage, n),
	}
}

func (gw *GoroutineWorkerPool) Start() {
	for i := 0; i < gw.workers; i++ {
		go gw.worker()
	}
}

func (gw *GoroutineWorkerPool) worker() {
	for {
		urlToCrawl, err := gw.repo.GetUnvisitedURL()
		if err != nil {
			// Handle error or break if no more URLs to process
			fmt.Println("error: gw.repo.GetUnvisitedURL:", err)
			break
		}

		content, err := gw.client.Get(urlToCrawl.URL)
		if err != nil {
			// Handle error
			fmt.Println("error: gw.client.Get:", err)
			continue
		}

		// Here, you can parse the content to extract new links
		// and add them to the repository or perform other tasks.
		fmt.Println("content: ", content)

		gw.repo.MarkAsVisited(urlToCrawl.URL)
	}
}

func (gw *GoroutineWorkerPool) AddJob(webPage *entities.WebPage) {
	gw.jobs <- webPage
}
