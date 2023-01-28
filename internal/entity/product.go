package entity

import uuid "github.com/satori/go.uuid"

type Product struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Price       int       `json:"price"`
	Currency    string    `json:"currency"`
	LeftInStock int       `json:"left_in_stock"`
}
