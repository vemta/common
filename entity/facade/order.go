package facade

type DB_Order struct {
<<<<<<< HEAD
	ID                 string  `db:"Id"`
	Customer           string  `db:"Customer"`
	Price              float64 `db:"Price"`
	PaymentMethod      string  `db:"Payment_method"`
	DiscountRaw        float64 `db:"DiscountRaw"`
=======
	ID                 string  `db:"id"`
	Customer           string  `db:"customer"`
	Price              float64 `db:"price"`
	PaymentMethod      int     `db:"payment_method"`
	DiscountRaw        float64 `db:"discount_raw"`
>>>>>>> 885aec38852b6cfc0ff91356ff8cb8a103e363ae
	DiscountPercentual float64 `db:"discount_percentual"`
	Status             int64   `db:"Status"`
}

type DB_OrderDetail struct {
<<<<<<< HEAD
	Order    string `db:"Order"`
	Item     string `db:"Item"`
	Quantity int    `db:"Quantity"`
=======
	Order    string `db:"order"`
	Item     string `db:"item"`
	Quantity int    `db:"quantity"`
>>>>>>> 885aec38852b6cfc0ff91356ff8cb8a103e363ae
}
