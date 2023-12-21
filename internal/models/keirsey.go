package models

type TestKeirsey struct {
	Res    string
	UserID string
	Date   string
}

func (t *TestKeirsey) Get() []string {
	return []string{"Keirsey", t.Res, t.UserID, t.Date}
}

// func (t *TestKeirsey)
