package models

type OrderLine struct {
	OrderHeaderId int     `json:"OrderHeaderId"`
	ItemId        int     `json:"ItemId"`
	Item          []Item  `json:"ItemId"`
	Amount        int     `json:"Amount"`
	LineNumber    int     `json:"LineNumber"`
	LinePrice     float64 `json:"LinePrice"`
}
