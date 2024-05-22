package models

type User struct {
	ID         string `json:"id"`
	Password   string `json:"password"`
	Lastname   string `json:"lastname"`
	Firstname  string `json:"firstname"`
	Patronymic string `json:"patronymic"`
	Email      string `json:"email"`
	Uni        string `json:"uni"`
	Syllabus   string `json:"syllabus"`
	Group      string `json:"group"`
	Age        string `json:"age"`
	City       string `json:"city"`
	Grade      string `json:"grade"`
}
