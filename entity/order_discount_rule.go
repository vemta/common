package entity

import (
	"time"
)

type OrderDiscountRule struct {
	ID                 string
	Name               string
	DiscountRaw        float64
	DiscountPercentual float64
	ApplyFirst         string // raw | percentual
	AboveValue         float64
	BellowValue        float64
	ValidFrom          time.Time
	ValidUntil         time.Time
}

func (d *OrderDiscountRule) TryApply(order *Order) float64 {

	if order.Price < d.AboveValue && d.AboveValue != -1 {
		return order.Price
	}

	if order.Price > d.BellowValue && d.BellowValue != -1 {
		return order.Price
	}

	if !d.ValidFrom.IsZero() && !time.Now().After(d.ValidFrom) {
		return order.Price
	}

	if !d.ValidUntil.IsZero() && !time.Now().Before(d.ValidUntil) {
		return order.Price
	}

	if d.ApplyFirst == "raw" {
		return (order.Price - d.DiscountRaw) * (1 - d.DiscountPercentual)
	}

	return (order.Price * (1 - d.DiscountPercentual)) - d.DiscountRaw
}
