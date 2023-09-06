package entity

import (
	"github.com/vemta/common/enum/orderstatus"
)

type Order struct {
	ID            string                  `json:"id"`
	Customer      *Customer               `json:"customer"`
	Items         *[]OrderEntry           `json:"items"`
	Price         float64                 `json:"price"`
	PaymentMethod int                     `json:"payment_method"`
	Status        orderstatus.OrderStatus `json:"status"`
}

type OrderEntry struct {
	Item     *Item `json:"item"`
	Quantity int64 `json:"quantity"`
}
