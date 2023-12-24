package models

type Res struct {
	Res    string
	UserID string
	Date   string
}

type Test struct {
	Test    string
	Results []Res
}
