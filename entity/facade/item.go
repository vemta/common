package facade

type DB_Item struct {
	ID          string `db:"id"`
	Title       string `db:"name"`
	Description string `db:"description"`
	IsGood      bool   `db:"is_good"`
	CreatedAt   string `db:"created_at"`
}

type DB_ItemValorizationTable struct {
	Item      string  `db:"item"`
	LastCost  float64 `db:"last_cost"`
	LastPrice float64 `db:"last_price"`
	Discount  float64 `db:"discount"`
}

type DB_ItemValorizationLog struct {
	Value     float64 `db:"price"`
	UpdatedAt string  `db:"updated_at"`
}

type DB_ItemTaxes struct {
	Name       string  `db:"name"`
	Percentual float64 `db:"percentual"`
	RawValue   float64 `db:"raw_value"`
}
