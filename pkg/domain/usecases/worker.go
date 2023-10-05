package usecases

import (
	"fmt"
	"github.com/scyanh/crawler/pkg/domain/entities"
	"github.com/scyanh/crawler/pkg/domain/interfaces"
	"golang.org/x/net/html"
	"strings"
	"sync"
)

type Worker struct {
	Repo       interfaces.IMemoryLinkRepository
	HTTPClient interfaces.IHTTPClient
}

func (w *Worker) Work(workerID int, wgWorkers, wgURLs *sync.WaitGroup, toVisitChan chan string, visitedChan chan entities.Link) {
	defer wgWorkers.Done()

	for url := range toVisitChan {
		if w.Repo.IsFirstVisit(url) {

			links, err := w.getLinks(workerID, url)
			if err != nil {
				fmt.Println("Error crawling:", url, "-", err)
			} else {

				for _, link := range links {
					//linkEntity2 := entities.Link{URL: link}
					//if !w.Repo.HasBeenVisited(linkEntity2) {
					wgURLs.Add(1) // Incrementa por cada nueva URL que agregues al canal
					toVisitChan <- link
					//}
				}
			}

			visitedLink := entities.Link{
				URL:   url,
				Links: links,
			}
			visitedChan <- visitedLink
		}

		wgURLs.Done() // Decrementa después de procesar la URL actual
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
					if w.isValidLink(a.Val) {
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

func (w *Worker) isValidLink(url string) bool {
	if strings.HasPrefix(url, "https://parserdigital.com") && url != "#" && url != "/" && url != "" {
		return true
	}

	return false
}
