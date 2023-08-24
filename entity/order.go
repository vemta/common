package entity

import (
	"github.com/vemta/common/enum"
)

type Order struct {
	ID                 string           `json:"id"`
	Customer           *User            `json:"customer"`
	Items              *[]OrderEntry    `json:"items"`
	Price              float64          `json:"price"`
	PaymentMethod      int              `json:"payment_method"`
	DiscountRaw        float64          `json:"discount_raw"`
	DiscountPercentual float64          `json:"discount_percentual"`
	Status             enum.OrderStatus `json:"status"`
}

type OrderEntry struct {
	Item     *Item `json:"item"`
	Quantity int64 `json:"quantity"`
}
