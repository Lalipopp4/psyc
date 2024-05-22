package models

type Review struct {
	Review   string `json:"review" db:"review"`
	UserID   string `json:"user_id" db:"user_id"`
	AuthorID string `json:"author_id" db:"author_id"`
	ResultID string `json:"result_id" db:"result_id"`
}
