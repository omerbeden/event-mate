package caching

type Cache interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	Exist(key string) (bool, error)
}
