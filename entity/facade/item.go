package facade

type DB_Item struct {
	ID          string `db:"ID"`
	Title       string `db:"Title"`
	Description string `db:"Description"`
	IsGood      bool   `db:"IsGood"`
	CreatedAt   string `db:"CreatedAt"`
}

type DB_ItemValuationTable struct {
	Item               string  `db:"Item"`
	LastCost           float64 `db:"LastCost"`
	LastPrice          float64 `db:"LastPrice"`
	DiscountRaw        float64 `db:"DiscountRaw"`
	DiscountPercentual float64 `db:"DiscountPercentual"`
}

type DB_ItemValuationLog struct {
	Price              float64 `db:"Price"`
	DiscountRaw        float64 `db:"DiscountRaw"`
	DiscountPercentual float64 `db:"DiscountPercentual"`
	UpdatedAt          string  `db:"UpdatedAt"`
}

type DB_ItemTaxes struct {
	Name       string  `db:"Name"`
	Percentual float64 `db:"Percentual"`
	RawValue   float64 `db:"RawValue"`
}
