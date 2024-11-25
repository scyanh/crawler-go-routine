package db

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	url = "https://example.com"
)

func TestNewInMemoryURLRepository(t *testing.T) {
	repo := NewInMemoryURLRepository()

	assert.NotNil(t, repo)
	assert.NotNil(t, repo.visitedURLs)
	assert.Empty(t, repo.visitedURLs)
}

func TestIsFirstVisit(t *testing.T) {
	repo := NewInMemoryURLRepository()

	// Test visiting the URL for the first time
	isFirstVisit := repo.IsFirstVisit(url)
	assert.True(t, isFirstVisit, "The URL should be marked as first visit")

	// Test when the URL is visited again
	isFirstVisit = repo.IsFirstVisit(url)
	assert.False(t, isFirstVisit, "The URL should have been visited before")
}
