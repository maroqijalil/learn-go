package models

type Item struct {
	ItemId      uint   `gorm:"primarykey;unique;autoIncrement" json:"lineItemId"`
	ItemCode    string `json:"itemCode"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderId     uint   `json:"orderId"`
}
