package controller

import (
	"app/models"
	"encoding/json"
	"log"
	"os"
	"sort"
	"time"
)

// 1. Order boyicha default holati time sort bolishi kerak. DESC
func (c *Controller) Task_1() []*models.Order {

	resp, _ := c.OrderGetList(&models.OrderGetListRequest{})
	orders := resp.Orders

	sort.Slice(orders, func(i, j int) bool {
		time1, _ := time.Parse("2006-01-02 15:04:05", orders[i].DateTime)
		time2, _ := time.Parse("2006-01-02 15:04:05", orders[j].DateTime)
		return time1.After(time2)
	})

	return orders
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

func (c *Controller) Task_3(ID string) []models.Task3Model {
	products, _ := readProduct("data/product.json")
	cards, _ := readCards("data/shop_cart.json")
	users, _ := readUser("data/user.json")

	userIdName := make(map[string]string)
	for _, u := range users {
		userIdName[u.Id] = u.Name
	}
	productsIdName := make(map[string]string)
	productIdPrice := make(map[string]int)
	for _, p := range products {
		productIdPrice[p.Id] = p.Price
		productsIdName[p.Id] = p.Name
	}

	type countTime struct {
		Count int
		Time  string
	}

	userMap := make(map[string]map[string]countTime)
	for _, u := range cards {
		if _, ok := userMap[u.UserId]; !ok {
			userMap[u.UserId] = make(map[string]countTime)
		}
		v := userMap[u.UserId][u.ProductId]
		v.Time = u.Time[:11]
		v.Count = u.Count
		userMap[u.UserId][u.ProductId] = v
	}

	var result []models.Task3Model

	for userId, innerMap := range userMap {
		for prodId, v := range innerMap {
			if ID == userId {
				result = append(result, models.Task3Model{
					UserName: userIdName[userId],
					Data: []models.Task3{
						{
							ProductName: productsIdName[prodId],
							Price:       productIdPrice[prodId],
							Count:       v.Count,
							Total:       (productIdPrice[prodId] * v.Count),
							Time:        v.Time,
						},
					},
				})
			}
		}
	}

	return result
}

// 4. User qancha pul mahsulot sotib olganligi haqida hisobot.
func (c *Controller) Task_4() {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
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
}

// 5. Productlarni Qancha sotilgan boyicha hisobot
func (c *Controller) Task_5() []models.Task5 {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	// shop card getlist method yo'q shu sababli o'zim datani oldim
	cards, _ := readCards("data/shop_cart.json")

	productIdName := make(map[string]string)
	for _, p := range products {
		productIdName[p.Id] = p.Name
	}

	prodIdTotalSell := make(map[string]int)
	for _, c := range cards {
		prodIdTotalSell[c.ProductId] += c.Count
	}
	var resp []models.Task5

	for id, count := range prodIdTotalSell {
		resp = append(resp, models.Task5{
			Name:  productIdName[id],
			Count: count,
		})
	}
	return resp
}

// 6. Top 10 ta sotilayotgan mahsulotlarni royxati.
func (c *Controller) Task_6() []models.Task5 {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	// shop card getlist method yo'q shu sababli o'zim datani oldim
	cards, _ := readCards("data/shop_cart.json")

	productIdName := make(map[string]string)
	for _, p := range products {
		productIdName[p.Id] = p.Name
	}
	prodIdTotalSell := make(map[string]int)
	for _, c := range cards {
		prodIdTotalSell[c.ProductId] += c.Count
	}

	var resp []models.Task5
	for id, count := range prodIdTotalSell {
		resp = append(resp, models.Task5{
			Name:  productIdName[id],
			Count: count,
		})
	}
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Count > resp[j].Count
	})
	return resp[:10]
}

// 7. TOP 10 ta Eng past sotilayotgan mahsulotlar royxati
func (c *Controller) Task_7() []models.Task5 {
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	products := productData.Products
	// shop card getlist method yo'q shu sababli o'zim datani oldim
	cards, _ := readCards("data/shop_cart.json")

	productIdName := make(map[string]string)
	for _, p := range products {
		productIdName[p.Id] = p.Name
	}

	prodIdTotalSell := make(map[string]int)
	for _, c := range cards {
		prodIdTotalSell[c.ProductId] += c.Count
	}
	var resp []models.Task5
	for id, count := range prodIdTotalSell {
		resp = append(resp, models.Task5{
			Name:  productIdName[id],
			Count: count,
		})
	}
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Count < resp[j].Count
	})
	return resp[:10]
}

//  9. Qaysi category larda qancha mahsulot sotilgan boyicha jadval
//     Name: Electronika Count: 12
func (c *Controller) Task_9() map[string]int {
	categoryData, _ := c.CategoryGetList(&models.CategoryGetListRequest{})
	productData, _ := c.ProductGetList(&models.ProductGetListRequest{})
	// shop card getlist method yo'q shu sababli o'zim datani oldim
	cards, _ := readCards("data/shop_cart.json")
	category := categoryData.Categorys
	products := productData.Products

	prodIdCategoryId := make(map[string]string)
	for _, p := range products {
		prodIdCategoryId[p.Id] = p.CategoryID
	}
	categoryName := make(map[string]string)
	for _, c := range category {
		categoryName[c.Id] = c.Name
	}

	categoryCount := make(map[string]int)
	for _, c := range cards {
		categoryCount[categoryName[prodIdCategoryId[c.ProductId]]] += c.Count
	}
	return categoryCount
}

// 10. Qaysi User eng Active xaridor. Bitta ma'lumot chiqsa yetarli.
func (c *Controller) Task_10() []models.Task5 {
	userData, _ := c.UserGetList(&models.UserGetListRequest{})
	users := userData.Users
	// shop card getlist method yo'q shu sababli o'zim datani oldim
	cards, _ := readCards("data/shop_cart.json")

	userIdName := make(map[string]string)
	for _, u := range users {
		userIdName[u.Id] = u.Name
	}

	userIdCount := make(map[string]int)
	for _, c := range cards {
		userIdCount[userIdName[c.UserId]] += c.Count
	}
	var resp []models.Task5
	for name, count := range userIdCount {
		resp = append(resp, models.Task5{
			Name:  name,
			Count: count,
		})
	}
	sort.Slice(resp, func(i, j int) bool {
		return resp[i].Count > resp[j].Count
	})
	return resp[:1]
}

//  11. Agar User 9 dan kop mahuslot sotib olgan bolsa,
//     1 tasi tekinga beriladi va 9 ta uchun pul hisoblanadi.
//     1 tasi eng arzon mahsulotni pulini hisoblamaysiz.
func (c *Controller) Task_11() int {
	resp, _ := c.OrderGetList(&models.OrderGetListRequest{})
	orders := resp.Orders

	totalSum := 0
	for _, order := range orders {
		orderSum := order.Sum
		orderCount := order.SumCount

		if orderCount > 9 {
			sort.SliceStable(order.OrderItems, func(i, j int) bool {
				return order.OrderItems[i].TotalPrice < order.OrderItems[j].TotalPrice
			})
			cheapestPrice := order.OrderItems[0].TotalPrice
			orderSum -= cheapestPrice
		}
		totalSum += orderSum
	}

	return totalSum

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
