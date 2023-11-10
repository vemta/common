package usecaseinput

import (
	"time"

	"github.com/vemta/common/entity"
	"github.com/vemta/common/enum"
)

type UpdateOrderStatusUsecaseInput struct {
	Order  string           `json:"order"`
	Status enum.OrderStatus `json:"status"`
}

type CreateDiscountRuleUsecaseInput struct {
	ID                 string    `json:"ID"`
	Name               string    `json:"name"`
	DiscountRaw        float64   `json:"discount_raw"`
	DiscountPercentual float64   `json:"discount_percentual"`
	ApplyFirst         string    `json:"apply_first"` // raw | percentual
	AboveValue         float64   `json:"above_value"`
	BellowValue        float64   `json:"bellow_value"`
	ValidFrom          time.Time `json:"valid_from"`
	ValidUntil         time.Time `json:"valid_until"`
	AutoApply          bool      `json:"auto_apply"`
	Items              []string  `json:"item,omitempty"`
	Code               string    `json:"code,omitempty"`
	Type               string    `json:"type"`
}

type CreateItemUsecaseInput struct {
	ID          string                `json:"id"`
	Title       string                `json:"title"`
	IsGood      bool                  `json:"is_good"`
	Description string                `json:"description"`
	Category    *entity.ItemCategory  `json:"category"`
	Valuation   *entity.ItemValuation `json:"valuation"`
}

type CreateOrderUsecaseInput struct {
	Customer      string               `json:"customer"`
	Items         *[]entity.OrderEntry `json:"items"`
	PaymentMethod int                  `json:"payment_method"`
}

type FindOrderFinalPriceUsecaseInput struct {
	DiscountCodes []string `json:"discount_codes"`
	Items         []string `json:"items"`
}

type ItemValorizationUpdateUsecaseInput struct {
	Item  string  `json:"item"`
	Cost  float64 `json:"cost"`
	Price float64 `json:"price"`
}
