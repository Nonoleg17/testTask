package entity

import uuid "github.com/satori/go.uuid"

type OrderProduct struct {
	OrderId       uuid.UUID `json:"order_id"`
	ProductId     uuid.UUID `json:"product_id"`
	CountProducts int       `json:"count_products"`
}
