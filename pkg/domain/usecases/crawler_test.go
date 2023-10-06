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
	repoMock := interfaces.NewMockIMemoryLinkRepository(ctrl)
	httpClientMock := interfaces.NewMockIHTTPClient(ctrl)

	setupMocksForURL(repoMock, "https://parserdigital.com")
	setupMocksForURL(repoMock, "https://parserdigital.com/about/")

	httpClientMock.EXPECT().Get("https://parserdigital.com").Return(htmlHome, nil).AnyTimes()
	httpClientMock.EXPECT().Get("https://parserdigital.com/about/").Return(htmlAbout, nil).AnyTimes()

	// Initialize the crawler
	numWorkers := 4
	crawler := NewCrawler(repoMock, httpClientMock, numWorkers)

	// Execute the crawler
	startLink := entities.Link{URL: "https://parserdigital.com"}
	crawler.Crawl(startLink)

	firstVisit := repoMock.IsFirstVisit(startLink.URL)
	assert.False(t, firstVisit, "The URL should have been visited before.")

}

func setupMocksForURL(repoMock *interfaces.MockIMemoryLinkRepository, url string) {
	callCount := 0
	repoMock.EXPECT().IsFirstVisit(url).DoAndReturn(func(url interface{}) bool {
		callCount++
		if callCount == 1 {
			return true
		}
		return false
	}).AnyTimes()
}
