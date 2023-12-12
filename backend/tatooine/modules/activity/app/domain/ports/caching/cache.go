package caching

type Cache interface {
	Set(key string, value any) error
	Get(key string) (any, error)
	AddMember(key string, members ...any) error
	GetMembers(key string) ([]string, error)
}
