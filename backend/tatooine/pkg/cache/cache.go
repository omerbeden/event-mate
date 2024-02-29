package cache

import "context"

type Cache interface {
	Set(ctx context.Context, key string, value any) error
	Get(ctx context.Context, key string) (any, error)
	AddMember(ctx context.Context, key string, members ...any) error
	GetMembers(ctx context.Context, key string) ([]string, error)
	Delete(ctx context.Context, key string) error
	Close() error
}
