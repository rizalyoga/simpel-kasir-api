package models

type Product struct {
	ID    int    `json:"id" db:"id"`
	Name  string `json:"name" db:"name"`
	Price int    `json:"price" db:"price"`
	Stock int    `json:"stock" db:"stock"`
}
