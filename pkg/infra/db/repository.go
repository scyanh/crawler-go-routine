package db

import (
	"github.com/scyanh/crawler/pkg/domain/entities"
	"sync"
)

type IMemoryItemRepository interface {
	HasBeenVisited(url entities.Item) bool
	MarkAsVisited(url entities.Item)
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

func (r *MemoryURLRepository) HasBeenVisited(item entities.Item) bool {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.visitedURLs[item.URL]
}

func (r *MemoryURLRepository) MarkAsVisited(item entities.Item) {
	r.mu.Lock()
	r.visitedURLs[item.URL] = true
	r.mu.Unlock()
}
