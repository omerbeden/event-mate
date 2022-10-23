package caching

type Cache interface {
	AddToCache(key string, value interface{}) error
	GetFromCache(key string) (interface{}, error)
	UpdateCache(key string, value interface{}) error
	Exist(key string) bool
}
