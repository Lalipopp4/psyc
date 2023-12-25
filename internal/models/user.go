package models

type User struct {
	ID       string
	Password string
	Info
}

type Info struct {
	Lastname   string
	Firstname  string
	Patronymic string
	Email      string
	Uni        string
	Syllabus   string
	Group      string
	Age        string
	City       string
	Grade      string
}
