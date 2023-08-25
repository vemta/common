package facade

type DB_Order struct {
	ID                 string  `db:"Id"`
	Customer           string  `db:"Customer"`
	Price              float64 `db:"Price"`
	PaymentMethod      string  `db:"Payment_method"`
	DiscountRaw        float64 `db:"DiscountRaw"`
	DiscountPercentual float64 `db:"discount_percentual"`
	Status             int64   `db:"Status"`
}

type DB_OrderDetail struct {
	Order    string `db:"Order"`
	Item     string `db:"Item"`
	Quantity int    `db:"Quantity"`
}
