package models

type User struct {
	Id           int    `json:"Id"`
	Username     string `Json:"Username"`
	PasswordHash string `Json:"PasswordHash"`
	PasswordSalt string `Json:"PasswordSalt"`
}
