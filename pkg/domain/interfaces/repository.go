package interfaces

// IMemoryLinkRepository is an interface for memory link repositories.
type IMemoryLinkRepository interface {
	IsFirstVisit(url string) bool
}
