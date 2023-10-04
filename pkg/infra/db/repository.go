package db

import (
	"errors"
	"github.com/scyanh/crawler/pkg/domain/entities"
	"sync"
)

type InMemoryURLRepository struct {
	urls  map[string]entities.WebPage
	mutex sync.RWMutex // Para garantizar la concurrencia segura al modificar el mapa.
}

func NewInMemoryURLRepository() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		urls: make(map[string]entities.WebPage),
	}
}

func (r *InMemoryURLRepository) AddURL(url entities.WebPage) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	// Si la URL ya está en el repositorio, no la sobreescriba.
	if _, exists := r.urls[url.URL]; !exists {
		r.urls[url.URL] = url
	}
}

func (r *InMemoryURLRepository) GetUnvisitedURL() (*entities.WebPage, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	for _, page := range r.urls {
		if !page.Visited {
			return &page, nil
		}
	}

	return nil, errors.New("no unvisited URLs found")
}

func (r *InMemoryURLRepository) MarkAsVisited(url string) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	page, exists := r.urls[url]
	if exists {
		page.Visited = true
		r.urls[url] = page
	}
}
