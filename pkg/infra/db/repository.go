package db

import (
	"sync"
)

type MemoryURLRepository struct {
	visitedURLs map[string]bool
	mu          sync.Mutex
}

// NewInMemoryURLRepository creates a new instance of MemoryURLRepository.
func NewInMemoryURLRepository() *MemoryURLRepository {
	return &MemoryURLRepository{
		visitedURLs: make(map[string]bool),
	}
}

// IsFirstVisit checks if the given URL has been visited before. If it's the first visit,
// it marks the URL as visited and returns true. Otherwise, returns false.
func (r *MemoryURLRepository) IsFirstVisit(url string) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.visitedURLs[url] {
		return false
	}

	r.visitedURLs[url] = true
	return true
}
