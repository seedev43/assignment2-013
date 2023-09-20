package models

import "time"

type Order struct {
	ID           uint   `gorm:"primaryKey"`
	CustomerName string `json:"customer_name" gorm:"column:customer_name;type:text;not null"`
	Items        []Item
	OrderedAt    string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
