package tgbot

import "context"

type PsycBot interface {
	Run(ctx context.Context)
}
