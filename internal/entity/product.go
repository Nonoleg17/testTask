package entity

type Product struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Currency    string `json:"currency"`
	LeftInStock int    `json:"left_in_stock"`
}
