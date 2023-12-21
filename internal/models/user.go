package models

type User struct {
	ID string
	Info
}

type Info struct {
	Lastname   string
	Firstname  string
	Patronymic string
	Email      string
	Password   string
	Uni        string
	Syllabus   string
	Group      string
	Age        int8
}
