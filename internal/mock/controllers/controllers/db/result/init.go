package result

import "database/sql"

type resultRepository struct {
}

func New(db *sql.DB) Repository {
	return &resultRepository{}
}
