package models

type Category struct {
	Id		string	`form:"id" json:"id"`
	Nama	string	`form:"nama" json:"nama"`
}

type CategoryBody struct {
	Nama	string	`form:"nama" json:"nama"`	
}