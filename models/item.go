package models

import "gorm.io/gorm"

type Item struct {
	gorm.Model
	ItemCode    string `json:"item_code" gorm:"column:item_code;unique;not null"`
	Description string `json:"description" gorm:"type:text"`
	Quantity    int    `json:"quantity" gorm:"not null"`
	OrderID     uint   `json:"order_id"`
}
