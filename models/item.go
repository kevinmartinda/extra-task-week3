package models

type Item struct {
	Id			int		`form:"id" json:"id"`
	ItemId		string	`form:"item_id" json:"item_id"`
	Nama		string	`form:"nama" json:"nama"`
}

type ItemBody struct {
	ItemId		string	`form:"item_id" json:"item_id"`
	CategoryId	int		`form:"kategori_id" json:"kategori_id"`
}