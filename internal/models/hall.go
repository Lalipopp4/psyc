package models

type TestHall struct {
	Res    string
	UserID string
	Date   string
}

func (t *TestHall) Get() []string {
	return []string{t.Res, t.UserID, t.Date}
}
