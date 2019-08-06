package models

type CategoryResponse struct {
	Status	int 	`json:"status"`
	Message string 	`json:"message"`
	Data	[]Category
}

type ItemResponse struct {
	Status	int 	`json:"status"`
	Message string 	`json:"message"`
	Data	[]Item
}