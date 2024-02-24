package user

import (
	"database/sql"
)

type userRepository struct {
}

func New(db *sql.DB) Repository {
	return &userRepository{}
}
