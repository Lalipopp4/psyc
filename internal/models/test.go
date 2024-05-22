package models

type Result struct {
	Test   string `json:"test"`
	Res    string `json:"res"`
	UserID string `json:"user_id"`
	Date   string `json:"date"`
}

type Test struct {
	Test      string          `json:"name,omitempty"`
	UserID    string          `json:"user_id,omitempty"`
	Result    map[uint]string `json:"result,omitempty"`
	Questions []Question      `json:"questions,omitempty"`
}

type Question struct {
	Code        uint   `db:"code" json:"code"`
	Question    string `db:"question" json:"question"`
	Type        uint8  `db:"type" json:"type"`
	Answers     string `db:"answers" json:"answers,omitempty"`
	RightAnswer string `db:"right_answer" json:"right_answer"`
}
