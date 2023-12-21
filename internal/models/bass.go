package models

type TestBass struct {
	Res    string
	UserID string
	Date   string
}

func (t *TestBass) Get() []string {
	return []string{t.Res, t.UserID, t.Date}
}
