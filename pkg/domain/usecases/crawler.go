package usecases

import (
	"fmt"
	"github.com/scyanh/crawler/pkg/domain/entities"
	"github.com/scyanh/crawler/pkg/domain/interfaces"
	"sync"
)

type Crawler struct {
	Repo       interfaces.IMemoryLinkRepository
	WorkerPool []*Worker
}

func NewCrawler(repo interfaces.IMemoryLinkRepository, httpClient interfaces.IHTTPClient, numWorkers int) *Crawler {
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

func (c *Crawler) Crawl(startItem entities.Link) {
	toVisitChan := make(chan string, 100)
	visitedChan := make(chan entities.Link)

	go c.startWorkers(startItem, toVisitChan, visitedChan)

	for link := range visitedChan {
		fmt.Printf(link.String())
	}
}

func (c *Crawler) startWorkers(startItem entities.Link, toVisitChan chan string, visitedChan chan entities.Link) {
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
