package entity

type User struct {
	ID         int    `json:"id"`
	Firstname  string `json:"firstname"`
	Surname    string `json:"surname"`
	Middlename string `json:"middlename"`
	Fio        string `json:"fio"`
	Sex        string `json:"sex"`
	Age        int    `json:"age"`
}
