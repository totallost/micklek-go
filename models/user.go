package models

type user struct {
	Id           int    `json:"Id"`
	Username     string `Json:"Username"`
	PasswordHash string `Json:"PasswordHash"`
	PasswordSalt string `Json:"PasswordSalt"`
}
