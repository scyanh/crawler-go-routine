package db

import (
	"github.com/scyanh/crawler/pkg/domain/entities"
	"sync"
)

type IMemoryURLRepository interface {
	HasBeenVisited(url entities.URL) bool
	MarkAsVisited(url entities.URL)
}

type MemoryURLRepository struct {
	visitedURLs map[string]bool
	mu          sync.Mutex
}

func NewInMemoryURLRepository() *MemoryURLRepository {
	return &MemoryURLRepository{
		visitedURLs: make(map[string]bool),
	}
}

func (r *MemoryURLRepository) HasBeenVisited(url entities.URL) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.visitedURLs[url.Value]
}

func (r *MemoryURLRepository) MarkAsVisited(url entities.URL) {
	r.mu.Lock()
	r.visitedURLs[url.Value] = true
	r.mu.Unlock()
}
