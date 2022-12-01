package entities

import "time"

type Order struct {
	ID                  uint      `json:"id"`
	BuyerID             uint      `json:"buyer_id"`
	SellerID            uint      `json:"seller_id"`
	ProductID           uint      `json:"product_id"`
	DeliverySource      string    `json:"delivery_source"`
	DeliveryDestination string    `json:"delivery_destination"`
	Items               string    `json:"items"`
	Price               float64   `json:"price"`
	Quantity            uint      `json:"quantity"`
	TotalPrice          float64   `json:"total_price"`
	Status              string    `json:"status"`
	CreatedAt           time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt           time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
