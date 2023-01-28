package entity

import uuid "github.com/satori/go.uuid"

type User struct {
	ID         uuid.UUID `json:"id"`
	Firstname  string    `json:"firstname"`
	Surname    string    `json:"surname"`
	Middlename string    `json:"middlename"`
	//Fio        string `json:"fio"`
	Sex string `json:"sex"`
	Age int    `json:"age"`
}
