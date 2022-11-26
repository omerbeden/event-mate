package caching

type Cache interface {
	AddToCache(key string, value interface{}) error
	GetFromCache(key string) (interface{}, error)
	Exist(key string) (bool, error)
}
