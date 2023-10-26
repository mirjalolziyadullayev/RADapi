package models

type User struct {
	Id int
	Firstname string
	Lastname string
	EmailUsername string
	Password string
}

type UserContent struct {
	User User
	Notes []Notes
}