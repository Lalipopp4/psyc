package user

import (
	"database/sql"
)

type userRepository struct {
	cur *sql.DB
}

func New(db *sql.DB) Repository {
	return &userRepository{cur: db}
}
