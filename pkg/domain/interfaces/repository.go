package interfaces

type IMemoryLinkRepository interface {
	IsFirstVisit(url string) bool
}
