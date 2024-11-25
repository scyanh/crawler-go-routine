package usecases

import (
	"github.com/scyanh/crawler/domain/entities"
	"github.com/scyanh/crawler/domain/interfaces"
	"golang.org/x/net/html"
	"log"
	"strings"
	"sync"
)

const validDomain = "https://parserdigital.com"

type Worker struct {
	Repo       interfaces.IMemoryLinkRepository
	HTTPClient interfaces.IHTTPClient
}

// Work is the main function of the worker.
func (w *Worker) Work(workerID int, wgWorkers, wgURLs *sync.WaitGroup, toVisitChan chan string, visitedChan chan entities.Link) {
	defer wgWorkers.Done()

	for url := range toVisitChan {
		if w.Repo.IsFirstVisit(url) {
			visitedLink := entities.Link{URL: url}

			links, err := w.getLinks(workerID, url)
			if err != nil {
				log.Println("Error crawling:", url, "-", err)
				visitedLink.Error = true
			} else {
				for _, link := range links {
					wgURLs.Add(1)
					toVisitChan <- link
				}
			}

			visitedLink.Links = links
			visitedChan <- visitedLink
		}

		// Decrease the counter of pending URLs to visit
		wgURLs.Done()
	}
}

// getLinks returns all the links found in the HTML of the given URL.
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
					if w.isValidLink(a.Val) {
						//log.Printf("worker=%d valid link: %s \n", workerID, a.Val)
						links = append(links, a.Val)
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

// isValidLink returns true if the given URL is valid to visit based on the validDomain.
func (w *Worker) isValidLink(url string) bool {
	if strings.HasPrefix(url, validDomain) {
		return true
	}

	return false
}
