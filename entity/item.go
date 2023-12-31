package entity

import (
	"time"
)

type Item struct {
	ID          string         `json:"id"`
	Title       string         `json:"title"`
	IsGood      bool           `json:"is_good"`
	Description string         `json:"description"`
	Category    *ItemCategory  `json:"category"`
	Valuation   *ItemValuation `json:"valuation"`
}

type ItemValuation struct {
	Item         string             `json:"item"`
	LastCost     float64            `json:"last_cost"`
	LastPrice    float64            `json:"last_price"`
	UpdatedAt    time.Time          `json:"updated_at"`
	PriceHistory []ItemValuationLog `json:"price_history"`
	CostHistory  []ItemValuationLog `json:"cost_history"`
}

type ItemValuationLog struct {
	Value     float64   `json:"price"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Tax struct {
	Name       string  `json:"name"`
	Percentual float64 `json:"percentual,omitempty"`
	RawValue   float64 `json:"raw_value,omitempty"`
}
