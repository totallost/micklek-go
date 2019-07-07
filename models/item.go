package models

type Item struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	Price           float64 `json:"Price"`
	PhotoUrl        string  `json:"photoUrl"`
	Description     string  `json:"description"`
	IsActive        bool    `json:"isActive"`
	PhotoPublicName string  `json:"photoPublicName"`
}
