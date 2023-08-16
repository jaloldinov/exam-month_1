package models

type ShopCard struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}

type Task3 struct {
	ProductName string
	Price       int
	Count       int
	Total       int
	Time        string
}

type Task3Model struct {
	UserName string
	Data     []Task3
}
