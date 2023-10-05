package interfaces

import "github.com/scyanh/crawler/pkg/domain/entities"

type IMemoryItemRepository interface {
	HasBeenVisited(url entities.Item) bool
	MarkAsVisited(url entities.Item)
}
