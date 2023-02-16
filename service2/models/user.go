package models

type User struct {
	Id       int
	Email    string
	Salt     string
	Password string
}
