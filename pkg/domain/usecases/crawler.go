package usecases

import (
	"fmt"
	"github.com/scyanh/crawler/pkg/domain/entities"
	"github.com/scyanh/crawler/pkg/infra/db"
	"github.com/scyanh/crawler/pkg/infra/http"
	"sync"
)

type Crawler struct {
	Repo       db.IMemoryItemRepository
	WorkerPool []*Worker
}

func NewCrawler(repo db.IMemoryItemRepository, httpClient http.IHTTPClient, numWorkers int) *Crawler {
	workers := make([]*Worker, numWorkers)
	for i := 0; i < numWorkers; i++ {
		workers[i] = &Worker{
			Repo:       repo,
			HTTPClient: httpClient,
		}
	}
	return &Crawler{
		Repo:       repo,
		WorkerPool: workers,
	}
}

func (c *Crawler) Crawl(startItem entities.Item) {
	toVisitChan := make(chan string, 100)
	visitedChan := make(chan string)

	go c.startWorkers(startItem, toVisitChan, visitedChan)

	for url := range visitedChan {
		fmt.Println("visited URL:", url)
	}
}

func (c *Crawler) startWorkers(startItem entities.Item, toVisitChan, visitedChan chan string) {
	var wgWorkers sync.WaitGroup
	var wgURLs sync.WaitGroup

	for i, worker := range c.WorkerPool {
		wgWorkers.Add(1)
		go worker.Work(i, &wgWorkers, &wgURLs, toVisitChan, visitedChan)
	}

	toVisitChan <- startItem.URL
	wgURLs.Add(1)

	wgURLs.Wait()
	close(toVisitChan)
	wgWorkers.Wait()
	close(visitedChan)
}
