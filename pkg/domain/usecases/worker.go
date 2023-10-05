package usecases

import (
	"fmt"
	"github.com/scyanh/crawler/pkg/infra/db"
	"github.com/scyanh/crawler/pkg/infra/http"
	"golang.org/x/net/html"
	"strings"
	"sync"
)

type Worker struct {
	Repo       db.IMemoryURLRepository
	HTTPClient http.IHTTPClient
}

var mu sync.Mutex

func (w *Worker) Work(workerID int, wgWorkers, wgURLs *sync.WaitGroup, toVisitChan, visitedChan chan string, visited map[string]bool) {
	defer wgWorkers.Done()

	for url := range toVisitChan {
		links, err := w.getLinks(workerID, url)
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

func (w *Worker) getLinks(workerID int, url string) ([]string, error) {
	content, err := w.HTTPClient.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(strings.NewReader(content))
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