package model

type GetFeedCommandResult struct {
	Events   *[]Event
	CacheHit bool
}
