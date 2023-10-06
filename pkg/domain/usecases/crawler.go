package usecases

import (
	"fmt"
	"github.com/scyanh/crawler/pkg/domain/entities"
	"github.com/scyanh/crawler/pkg/domain/interfaces"
	"sync"
)

// Crawler is the struct that represents the crawler.
type Crawler struct {
	Repo       interfaces.IMemoryLinkRepository
	WorkerPool []*Worker
}

// NewCrawler returns a new Crawler instance.
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

// Crawl starts the crawling process from the given URL.
func (c *Crawler) Crawl(startItem entities.Link) {
	toVisitChan := make(chan string, 1000)
	visitedChan := make(chan entities.Link, 1000)

	go c.startWorkers(startItem, toVisitChan, visitedChan)

	var links []entities.Link

	for link := range visitedChan {
		fmt.Printf(link.String())
		links = append(links, link)
	}

	fmt.Printf("Total links: %d \n", len(links))
}

// startWorkers starts the workers and closes the channels when all the URLs have been visited.
func (c *Crawler) startWorkers(startItem entities.Link, toVisitChan chan string, visitedChan chan entities.Link) {
	var wgWorkers sync.WaitGroup
	var wgURLs sync.WaitGroup

	for i, worker := range c.WorkerPool {
		wgWorkers.Add(1)
		go worker.Work(i, &wgWorkers, &wgURLs, toVisitChan, visitedChan)
	}

	wgURLs.Add(1)
	toVisitChan <- startItem.URL

	// Wait until all the URLs have been visited
	go func() {
		wgURLs.Wait()
		close(toVisitChan)
	}()

	// Wait until all the workers have finished
	wgWorkers.Wait()
	close(visitedChan)
}
