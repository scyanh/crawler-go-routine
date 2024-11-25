package usecases

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/scyanh/crawler/domain/entities"
	"github.com/scyanh/crawler/domain/interfaces"
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestWorker_Work(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Setup Mocks
	repoMock := interfaces.NewMockIMemoryLinkRepository(ctrl)
	httpClientMock := interfaces.NewMockIHTTPClient(ctrl)

	httpClientMock.EXPECT().Get("https://parserdigital.com").Return("", errors.New("error")).AnyTimes()
	worker := &Worker{Repo: repoMock, HTTPClient: httpClientMock}

	setupMocksForURL(repoMock, "https://parserdigital.com")

	// Initialize channels and wait groups
	toVisitChan := make(chan string, 1000)
	visitedChan := make(chan entities.Link)
	wgWorkers := &sync.WaitGroup{}
	wgURLs := &sync.WaitGroup{}

	wgWorkers.Add(1)
	wgURLs.Add(1)
	go worker.Work(1, wgWorkers, wgURLs, toVisitChan, visitedChan)

	toVisitChan <- "https://parserdigital.com"
	close(toVisitChan)

	visitedLink := <-visitedChan

	assert.Equal(t, "https://parserdigital.com", visitedLink.URL)
	assert.True(t, visitedLink.Error)
}

func TestWorker_getLinks_ErrorFetching(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Setup Mocks
	repoMock := interfaces.NewMockIMemoryLinkRepository(ctrl)
	httpClientMock := interfaces.NewMockIHTTPClient(ctrl)

	httpClientMock.EXPECT().Get("https://parserdigital.com").Return("", errors.New("error")).AnyTimes()

	worker := &Worker{Repo: repoMock, HTTPClient: httpClientMock}
	links, err := worker.getLinks(1, "https://parserdigital.com")

	assert.Error(t, err)
	assert.Equal(t, 0, len(links))
}

func TestWorker_isValidLink(t *testing.T) {
	worker := &Worker{}
	assert.True(t, worker.isValidLink("https://parserdigital.com/link1"))
	assert.False(t, worker.isValidLink("https://notparserdigital.com/link1"))
}
