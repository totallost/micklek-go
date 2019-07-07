package dtos

type ClientInfoDto struct {
	ID           int    `Json:"ID"`
	FirstName    string `json:"FirstName"`
	SureName     string `json:"SureName"`
	MobileNumber string `json:"MobileNumber"`
	Email        string `json:"Email"`
	Address      string `json:"Address"`
	DateReady    string `json:"DateReady"`
	Notes        string `json:"Notes"`
	Status       int    `Json:"Status"`
}
