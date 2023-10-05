package db

import (
	"github.com/scyanh/crawler/pkg/domain/entities"
	"sync"
)

type MemoryURLRepository struct {
	visitedURLs map[string]bool
	mu          sync.Mutex
}

func NewInMemoryURLRepository() *MemoryURLRepository {
	return &MemoryURLRepository{
		visitedURLs: make(map[string]bool),
	}
}

func (r *MemoryURLRepository) HasBeenVisited(link entities.Link) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.visitedURLs[link.URL]
}

func (r *MemoryURLRepository) MarkAsVisited(link entities.Link) {
	r.mu.Lock()
	r.visitedURLs[link.URL] = true
	r.mu.Unlock()
}
