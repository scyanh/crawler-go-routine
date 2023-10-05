package usecases

import (
	"fmt"
	"github.com/scyanh/crawler/pkg/domain/entities"
	"github.com/scyanh/crawler/pkg/domain/interfaces"
	"runtime"
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

func (c *Crawler) startWorkers(startItem entities.Link, toVisitChan chan string, visitedChan chan entities.Link) {
	var wgWorkers sync.WaitGroup
	var wgURLs sync.WaitGroup

	for i, worker := range c.WorkerPool {
		wgWorkers.Add(1)
		go worker.Work(i, &wgWorkers, &wgURLs, toVisitChan, visitedChan)
	}

	wgURLs.Add(1)
	toVisitChan <- startItem.URL

	fmt.Printf("Current 1 Goroutines: %d\n", runtime.NumGoroutine())
	// Goroutine para cerrar el canal después de que todas las URLs se hayan visitado
	go func() {
		fmt.Printf("Current 2 Goroutines: %d\n", runtime.NumGoroutine())
		wgURLs.Wait()
		fmt.Printf("Current 3 Goroutines: %d\n", runtime.NumGoroutine())
		close(toVisitChan)
	}()

	fmt.Printf("Current 4 Goroutines: %d\n", runtime.NumGoroutine())
	wgWorkers.Wait()
	fmt.Printf("Current 5 Goroutines: %d\n", runtime.NumGoroutine())
	close(visitedChan)
}
