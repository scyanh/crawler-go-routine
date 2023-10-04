package usecases

import (
	"github.com/scyanh/crawler/pkg/domain"
	"github.com/scyanh/crawler/pkg/domain/entities"
)

type CrawlerInteractor struct {
	repo    domain.Repository
	workers entities.WorkerPool
}

func NewCrawlerInteractor(r domain.Repository, w entities.WorkerPool) *CrawlerInteractor {
	return &CrawlerInteractor{repo: r, workers: w}
}
func (ci *CrawlerInteractor) Crawl(startURL string) {
	// Start workers
	ci.workers.Start()

	// Add the starting URL as an initial job
	initialWebPage := &entities.WebPage{URL: startURL, Visited: false}
	ci.repo.AddURL(*initialWebPage)
	ci.workers.AddJob(initialWebPage)

	// Now, you could keep checking for unvisited URLs and keep adding to workers.
	// However, remember, since the workers themselves would be finding new URLs and
	// adding to the repo, you may want to design a way for the workers to signal
	// when there are truly no more URLs left to crawl to terminate the process gracefully.
}
