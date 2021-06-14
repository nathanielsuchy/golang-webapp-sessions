package models

type User struct {
	Id           int
	Email        string `orm:"unique"`
	PasswordHash string
}
