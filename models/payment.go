package models

import (
	"time"
)

type Payment struct {
	ID            uint      `gorm:"primaryKey"`
	OrderID       uint      `json:"order_id"`
	Provider      string    `json:"provider"` // e.g., Stripe, PayPal
	Amount        float64   `json:"amount"`
	Currency      string    `json:"currency"`
	TransactionID string    `json:"transaction_id" gorm:"unique"`
	Status        string    `json:"status"` // pending, success, failed
	CreatedAt     time.Time `json:"created_at"`
}
