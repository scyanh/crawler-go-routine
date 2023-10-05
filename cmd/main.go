package main

import (
	"fmt"
	"github.com/scyanh/crawler/pkg/domain/entities"
	"github.com/scyanh/crawler/pkg/domain/usecases"
	"github.com/scyanh/crawler/pkg/infra/db"
	"github.com/scyanh/crawler/pkg/infra/http"
	"runtime"
	"time"

	"sync"
)

var mu sync.Mutex

const (
	targetURL          = "https://parserdigital.com/"
	httpRequestTimeOut = 50 // value in seconds
	maxWorkers         = 10
)

func main() {
	// allow parallelism
	runtime.GOMAXPROCS(runtime.NumCPU())
	startTime := time.Now()

	// dependencies initialization
	repo := db.NewInMemoryURLRepository()
	httpClient := http.NewHTTPClient(httpRequestTimeOut * time.Second)
	crawler := usecases.NewCrawler(repo, httpClient, maxWorkers)

	// crawler execution
	startURL := entities.Item{URL: targetURL}
	crawler.Crawl(startURL)

	// print execution time
	duration := time.Since(startTime)
	fmt.Printf("Execution time: %s \n", duration)
}

/*
func crawler(startURL string, toVisitChan, visitedChan chan string, visited map[string]bool) {
	var wgWorkers sync.WaitGroup
	var wgURLs sync.WaitGroup

	numWorkers := 10
	for i := 0; i < numWorkers; i++ {
		wgWorkers.Add(1)
		go worker(i, &wgWorkers, &wgURLs, toVisitChan, visitedChan, visited)
	}

	wgURLs.Add(1)
	toVisitChan <- startURL

	wgURLs.Wait()
	close(toVisitChan)
	wgWorkers.Wait()
	close(visitedChan)
}

func worker(workerID int, wgWorkers, wgURLs *sync.WaitGroup, toVisitChan, visitedChan chan string, visited map[string]bool) {
	defer wgWorkers.Done()

	for url := range toVisitChan {
		links, err := getLinks(workerID, url)
		if err != nil {
			fmt.Println("Error crawling:", url, "-", err)
			wgURLs.Done()
			continue
		}

		mu.Lock()
		for _, link := range links {
			if _, seen := visited[link]; !seen {
				visited[link] = true
				wgURLs.Add(1)
				toVisitChan <- link
			}
		}
		visitedChan <- url
		mu.Unlock()

		wgURLs.Done()
	}
}

func getLinks(workerID int, url string) ([]string, error) {
	client := &http.Client{
		Timeout: 50 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {

					if strings.HasPrefix(a.Val, "https://parserdigital.com") && a.Val != "#" && a.Val != "/" && a.Val != "" {
						fmt.Printf("worker=%d a.Val filtered: %s \n", workerID, a.Val)
						links = append(links, a.Val)
					} else {
						fmt.Printf("worker=%d a.Val removed: %s \n", workerID, a.Val)
					}
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return links, nil
}
*/
