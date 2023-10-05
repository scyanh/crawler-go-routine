package usecases

import (
	"github.com/golang/mock/gomock"
	"github.com/scyanh/crawler/pkg/domain/entities"
	"github.com/scyanh/crawler/pkg/domain/interfaces"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func getHTMLFromFile(filename string) (string, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func TestCrawler_Crawl(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// test data
	htmlHome, err := getHTMLFromFile("../../../test_data/sample_home.html")
	if err != nil {
		t.Fatal(err)
	}
	htmlAbout, err := getHTMLFromFile("../../../test_data/sample_about.html")
	if err != nil {
		t.Fatal(err)
	}

	// Setup Mocks
	repoMock := interfaces.NewMockIMemoryItemRepository(ctrl)
	httpClientMock := interfaces.NewMockIHTTPClient(ctrl)

	//setupMocksForURL(t, repoMock, "https://parserdigital.com")
	//setupMocksForURL(t, repoMock, "https://parserdigital.com/about/")
	//setupMocksForURL(t, repoMock, "https://parserdigital.com/expertise/")

	itemHome := entities.Item{URL: "https://parserdigital.com"}
	/*
		callCount := 0
		repoMock.EXPECT().HasBeenVisited(itemHome).DoAndReturn(func(url interface{}) bool {
			callCount++
			if callCount == 1 {
				return false
			}
			return true
		}).AnyTimes()*/
	gomock.InOrder(
		repoMock.EXPECT().HasBeenVisited(itemHome).Return(false),
		repoMock.EXPECT().HasBeenVisited(itemHome).Return(true),
	)

	itemAbout := entities.Item{URL: "https://parserdigital.com/about/"}
	gomock.InOrder(
		repoMock.EXPECT().HasBeenVisited(itemAbout).Return(false),
		repoMock.EXPECT().HasBeenVisited(itemAbout).Return(true),
	)

	/*callCount3 := 0
	repoMock.EXPECT().HasBeenVisited("https://parserdigital.com/expertise/").DoAndReturn(func(url interface{}) bool {
		callCount3++
		if callCount3 == 1 {
			return false
		}
		return true
	}).AnyTimes()*/

	repoMock.EXPECT().MarkAsVisited(gomock.Any()).AnyTimes()
	httpClientMock.EXPECT().Get("https://parserdigital.com").Return(htmlHome, nil).AnyTimes()
	httpClientMock.EXPECT().Get("https://parserdigital.com/about/").Return(htmlAbout, nil).AnyTimes()

	// Initialize the crawler
	numWorkers := 4
	crawler := NewCrawler(repoMock, httpClientMock, numWorkers)

	// Execute the crawler
	startItem := entities.Item{URL: "https://parserdigital.com"}
	crawler.Crawl(startItem)

	visited := repoMock.HasBeenVisited(startItem)
	assert.True(t, visited, "The URL should be marked as visited")
}

func setupMocksForURL(t *testing.T, repoMock *interfaces.MockIMemoryItemRepository, url string) {
	callCount := 0
	repoMock.EXPECT().HasBeenVisited(url).DoAndReturn(func(url interface{}) bool {
		callCount++
		if callCount == 1 {
			return false
		}
		return true
	}).AnyTimes()
}
