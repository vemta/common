package entity

import (
	"github.com/vemta/common/enum"
)

type Order struct {
	ID            string             `json:"id"`
	Customer      *Customer          `json:"customer"`
	Items         *[]OrderEntry      `json:"items"`
	Price         float64            `json:"price"`
	PaymentMethod enum.PaymentMethod `json:"payment_method"`
	Status        enum.OrderStatus   `json:"status"`
}

type OrderEntry struct {
	Item               *Item   `json:"item"`
	DiscountRaw        float64 `json:"item_discount_raw"`
	DiscountPercentual float64 `json:"item_discount_percentual"`
	Quantity           int64   `json:"quantity"`
}
