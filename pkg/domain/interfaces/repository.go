package interfaces

// IMemoryLinkRepository is an interface for memory link repository.
type IMemoryLinkRepository interface {
	IsFirstVisit(url string) bool
}
