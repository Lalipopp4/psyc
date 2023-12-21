package result

import "database/sql"

type resultRepository struct {
	cur   *sql.DB
	tests []string
}

func New(db *sql.DB) Repository {
	return &resultRepository{cur: db}
}
