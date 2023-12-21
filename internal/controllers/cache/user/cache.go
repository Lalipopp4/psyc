package user

import "context"

type Cache interface {
	Check(ctx context.Context, key string) bool
	Add(ctx context.Context, key string) bool
}
