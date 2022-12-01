package entities

import "time"

type Product struct {
	ID          uint      `json:"id"`
	ProductName string    `json:"name"`
	Price       float64   `json:"price"`
	Description string    `json:"description"`
	SellerID    uint      `json:"seller_id"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
