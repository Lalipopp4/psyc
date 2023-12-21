package result

import "psyc/internal/controllers/db/result"

type resultService struct {
	repo result.Repository
}

func New(repo result.Repository) Service {
	return &resultService{repo: repo}
}
