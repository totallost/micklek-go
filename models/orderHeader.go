package models

type OrderHeader struct {
	Id              int         `json:"Id"`
	NumberOfItems   int         `json:"NumberOfItems`
	TotalPrice      float64     `json:TotalPrice`
	ClientFirstName string      `json:"ClientFirstName"`
	ClientSureName  string      `json:"ClientSureName"`
	ClientEmail     string      `json:"ClientEmail"`
	ClientCell      string      `json:"ClientCell"`
	Address         string      `json:"Address"`
	ClientRemarks   string      `json:"ClientRemarks"`
	OrderLines      []OrderLine `json:"OrderLines"`
	DateCreation    string      `json:"DateCreation"`
	DateTarget      string      `json:"DateTarget"`
	Status          []Status    `json:"Status"`
	StatusId        int         `json:StatusId`
}
