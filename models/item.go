package models

import "time"

type Item struct {
	ID          uint   `gorm:"primaryKey"`
	ItemCode    string `json:"item_code" gorm:"column:item_code;not null;unique"`
	Description string `gorm:"type:text"`
	Quantity    int    `gorm:"not null"`
	OrderID     uint
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
