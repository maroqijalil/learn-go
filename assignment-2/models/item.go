package models

type Item struct {
	ItemId      uint `gorm:"primarykey"`
	ItemCode    string
	Description string
	Quantity    uint
	OrderId     uint
}
