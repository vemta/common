package entity

import (
	"time"

	"github.com/google/uuid"
)

type Item struct {
	ID          uuid.UUID `json:"id" db:"id"`
	IsGood      bool      `json:"is_good" db:"is_good"`
	Description string    `json:"description" db:"description"`
}

type ItemValorizationTable struct {
	Item         uuid.UUID          `json:"item" db:"item"`
	Discount     float64            `json:"discount" db:"discount"`
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
