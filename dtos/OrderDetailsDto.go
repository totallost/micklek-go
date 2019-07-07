package dtos

type OrderDetailsDto struct {
	ClienDetails ClientInfoDto   `json:"clienDetails"`
	OrderDetails []OrderLinesDto `json:"orderDetails"`
}
