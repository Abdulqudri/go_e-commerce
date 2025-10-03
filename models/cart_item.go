package models

import (
	"time"

	"gorm.io/gorm"
)

type CartItem struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id"`
	ProductID uint           `json:"product_id"`
	Quantity  int            `json:"quantity"`
	Product   Product        `json:"product" gorm:"foreignKey:ProductID"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
