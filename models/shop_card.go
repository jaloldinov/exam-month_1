package models

// "product_id": "ace982a1-d556-40b1-9089-45c7157624eb",
// "user_id": "48097741-22c9-4663-8796-3c9993d88ffe",
// "count": 5,
// "status": false,
// "time": "2022-03-08 13:27:39"

type ShopCard struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
	Count     int    `json:"count"`
	Status    bool   `json:"status"`
	Time      string `json:"time"`
}
