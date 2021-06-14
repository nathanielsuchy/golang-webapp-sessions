package models

type User struct {
	Id           int
	Email        string `orm:"unique"`
	PasswordHash string
	IsDisabled   bool `default:false`
	IsSuperUser  bool `default:false`
}
