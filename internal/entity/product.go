package entity

type Product struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Currency    string `json:"currency"`
	LeftInStock int    `json:"leftInStock"`
}
