package dto

type Product struct {
	ID          int64   `db:"id"`
	Name        string  `db:"name"`
	Description string  `db:"description"`
	Category    string  `db:"category"`
	Price       float64 `db:"price"`
	Is_seasonal bool    `db:"is_seasonal"`
	Quantity    int64   `db:"quantity"`
}
