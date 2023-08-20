package facade

type DB_Order struct {
	ID                 string  `db:"id"`
	Customer           string  `db:"customer"`
	Price              float64 `db:"price"`
	PaymentMethod      string  `db:"payment_method"`
	DiscountRaw        float64 `db:"discount_raw"`
	DiscountPercentual float64 `db:"discount_percentual"`
	Status             int64   `db:"status"`
}

type DB_OrderDetail struct {
	Order    string `db:"order"`
	Item     string `db:"item"`
	Quantity int    `db:"quantity"`
}
