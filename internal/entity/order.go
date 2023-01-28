package entity

import uuid "github.com/satori/go.uuid"

type Order struct {
	ID     uuid.UUID `json:"id"`
	UserId uuid.UUID `json:"userId"`
}
