package entity

type DiscountCode struct {
	DiscountRuleInterface
	*DiscountRule
	Code string
}
