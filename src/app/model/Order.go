package model

type Order struct{
	Model
	UserId uint `json:"user_id"`
	CategoryId uint `json:"category_id"`
	ProductId uint `json:"product_id"`
}

