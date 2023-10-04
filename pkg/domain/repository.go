package domain

import "github.com/scyanh/crawler/pkg/domain/entities"

type Repository interface {
	AddURL(url entities.WebPage)
	GetUnvisitedURL() (*entities.WebPage, error)
	MarkAsVisited(url string)
}
