package result

import (
	"database/sql"
	"psyc/internal/controllers/cache"
)

type resultRepository struct {
	cur   *sql.DB
	cache cache.Cache
}

func New(db *sql.DB, cache cache.Cache) Repository {
	return &resultRepository{cur: db, cache: cache}
}
