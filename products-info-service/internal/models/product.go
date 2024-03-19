package models

type Product struct {
	ID          int     `db:"id" json:"id"`
	Name        string  `db:"name" json:"name"`
	Description string  `db:"description" json:"description"`
	Photo       string  `db:"photo" json:"photo"`
	Price       float64 `db:"price" json:"price"`
	Quantity    int     `db:"quantity" json:"quantity"`
}

type ProductLightweight struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

type Quantity struct {
	ID    int     `db:"id" json:"id" redis:"id"`
	Name  string  `db:"name" json:"name" redis:"name"`
	Count int     `db:"quantity" json:"quantity" redis:"quantity"`
	Price float64 `db:"price" json:"price" redis:"price"`
}
