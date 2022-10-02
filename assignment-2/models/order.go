package models

import "time"

type Order struct {
	OrderId      uint      `gorm:"primarykey;unique;autoIncrement" json:"orderId"`
	CustomerName string    `json:"customerName"`
	OrderedAt    time.Time `json:"orderedAt"`
	Items        []Item    `gorm:"foreignKey:OrderId" json:"items"`
}
