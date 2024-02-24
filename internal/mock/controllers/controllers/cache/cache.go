package cache

import "context"

type Cache interface {
	Check(ctx context.Context, acl, key string) bool
	Add(ctx context.Context, key, value string) error
}
