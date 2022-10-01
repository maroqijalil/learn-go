package models

import "time"

type Order struct {
	OrderId      uint `gorm:"primaryKey"`
	CustomerName string
	OrderedAt    time.Time
	Items        []Item `gorm:"foreignKey:OrderId"`
}
