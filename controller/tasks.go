package controller

import (
	"app/models"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// 1. Order boyicha default holati time sort bolishi kerak. DESC
func (c *Controller) Task_1() {

}

// 2. Order Date boyicha filter qoyish kerak
// M: 2022-02-01 dan 2022-02-22 gacha
// from_date: 2022-02-22 to_date:  2022-02-22

// 3. User history chiqish kerak. Ya'ni sotib olgan mahsulotlari korsatish kerak
// M:
//     Order Name: David
//     1. Name: Qora futbolka Price: 98000 Count: 3 Total: 294000 Time: 2022-03-29 11:10:52
//     2. Name: Qora futbolka Price: 98000 Count: 3 Total: 294000 Time: 2022-03-29 11:10:52
//     3. Name: Qora futbolka Price: 98000 Count: 3 Total: 294000 Time: 2022-03-29 11:10:52
//     4. Name: Qora futbolka Price: 98000 Count: 3 Total: 294000 Time: 2022-03-29 11:10:52

//  4. User qancha pul mahsulot sotib olganligi haqida hisobot.
//     Name: David Total Buy Price: 2022000
func (c *Controller) Task_4() {
	products, _ := readProduct("data/product.json")
	cards, _ := readCards("data/shop_cart.json")
	users, _ := readUser("data/user.json")

	userIdName := make(map[string]string)
	for _, u := range users {
		userIdName[u.Id] = u.Name
	}

	productIdPrice := make(map[string]int)
	for _, p := range products {
		productIdPrice[p.Id] = p.Price
	}

	userIdTotalSumPrice := make(map[string]int)
	for _, u := range cards {
		userIdTotalSumPrice[u.UserId] += (productIdPrice[u.ProductId] * u.Count)
	}

	for id, sum := range userIdTotalSumPrice {
		//	Name: David Total Buy Price: 2022000
		fmt.Printf("Name: %s => Buy Price: %d so'm\n", userIdName[id], sum)
	}
}

//  5. Productlarni Qancha sotilgan boyicha hisobot
//     Name: Asus count: 2
func (c *Controller) Task_5() {
	products, _ := readProduct("data/product.json")
	cards, _ := readCards("data/shop_cart.json")

	productIdName := make(map[string]string)
	for _, p := range products {
		productIdName[p.Id] = p.Name
	}

	prodIdTotalSell := make(map[string]int)
	for _, c := range cards {
		prodIdTotalSell[c.ProductId] += c.Count
	}

	for id, count := range prodIdTotalSell {
		//     Name: Asus count: 2
		fmt.Printf("Name: %s - Count: %d\n", productIdName[id], count)
	}
}

// ============== READERS ===================
func readCards(data string) ([]models.ShopCard, error) {
	var cards []models.ShopCard

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &cards)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return cards, nil
}

func readUser(data string) ([]models.User, error) {
	var users []models.User

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &users)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return users, nil
}

func readProduct(data string) ([]models.Product, error) {
	var products []models.Product

	d, err := os.ReadFile(data)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}
	err = json.Unmarshal(d, &products)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}
	return products, nil
}
