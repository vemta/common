package entity

import "time"

type DiscountRule struct {
	ID                 string
	Name               string
	DiscountRaw        float64
	DiscountPercentual float64
	ApplyFirst         string // raw | percentual
	AboveValue         float64
	BellowValue        float64
	ValidFrom          time.Time
	ValidUntil         time.Time
	Type               rune // (O)rder, (I)tem, (C)ode
}
