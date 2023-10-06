package main

import (
	"fmt"
	"github.com/scyanh/crawler/pkg/domain/entities"
	"github.com/scyanh/crawler/pkg/domain/usecases"
	"github.com/scyanh/crawler/pkg/infra/db"
	"github.com/scyanh/crawler/pkg/infra/http"
	"runtime"
	"time"
)

const (
	targetURL          = "https://parserdigital.com/"
	httpRequestTimeOut = 10 * time.Second
	maxWorkers         = 20
)

func main() {
	// Allow parallelism
	runtime.GOMAXPROCS(runtime.NumCPU())
	startTime := time.Now()

	// Dependencies initialization
	repo := db.NewInMemoryURLRepository()
	httpClient := http.NewHTTPClient(httpRequestTimeOut)
	crawler := usecases.NewCrawler(repo, httpClient, maxWorkers)

	// Crawler execution
	startURL := entities.Link{URL: targetURL}
	crawler.Crawl(startURL)

	// Execution time
	duration := time.Since(startTime)
	fmt.Printf("Execution time: %s \n", duration)
}
