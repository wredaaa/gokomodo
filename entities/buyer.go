package entities

import "time"

type Buyer struct {
	ID              uint      `json:"id"`
	Email           string    `json:"email"`
	Name            string    `json:"name"`
	Password        string    `json:"password"`
	ShippingAddress string    `json:"shipping_address"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
