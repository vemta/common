package entity

import (
	"time"

	"golang.org/x/exp/slices"
)

type ItemDiscountRule struct {
	DiscountRuleInterface
	*DiscountRule
	Items []string
}

func (d *ItemDiscountRule) TryApply(item *Item) (bool, float64) {

	if !slices.Contains(d.Items, item.ID) {
		return false, item.Valuation.LastPrice
	}

	if item.Valuation.LastPrice < d.AboveValue && d.AboveValue != -1 {
		return false, item.Valuation.LastPrice
	}

	if item.Valuation.LastPrice > d.BellowValue && d.BellowValue != -1 {
		return false, item.Valuation.LastPrice
	}

	if !d.ValidFrom.IsZero() && !time.Now().After(d.ValidFrom) {
		return false, item.Valuation.LastPrice
	}

	if !d.ValidUntil.IsZero() && !time.Now().Before(d.ValidUntil) {
		return false, item.Valuation.LastPrice
	}

	if d.ApplyFirst == "raw" {
		return true, (item.Valuation.LastPrice - d.DiscountRaw) * (1 - d.DiscountPercentual)
	}

	return true, (item.Valuation.LastPrice * (1 - d.DiscountPercentual)) - d.DiscountRaw
}
