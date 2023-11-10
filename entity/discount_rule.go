package entity

import "time"

type DiscountRule struct {
	ID                 string
	Code               string
	ApplyFirst         string // raw | percentual
	AutoApply          bool
	Name               string
	DiscountRaw        float64
	DiscountPercentual float64
	AboveValue         float64
	BellowValue        float64
	ValidFrom          time.Time
	ValidUntil         time.Time
	Type               rune // (O)rder, (I)tem, (C)ode
}

type DiscountRuleInterface interface{}
