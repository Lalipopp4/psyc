package result

import (
	"context"
)

type Repository interface {
	Get(ctx context.Context, arg any) (any, error)
	Add(ctx context.Context, data any) error
}
