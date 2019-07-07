package dtos

import (
	"models"
)

type OrderLinesDto struct {
	OrderID    int         `Json:"OrderId"`
	ItemID     int         `json:"ItemID"`
	Item       models.Item `json:"Item"`
	Amount     int         `json:"Amount"`
	LineNumber int         `json:"LineNumber"`
	LinePrice  float64     `json:"LinePrice"`
}
