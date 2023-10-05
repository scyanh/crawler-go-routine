package interfaces

import "github.com/scyanh/crawler/pkg/domain/entities"

type IMemoryLinkRepository interface {
	HasBeenVisited(link entities.Link) bool
	MarkAsVisited(link entities.Link)
}
