package main

import (
	"app/config"
	"app/controller"
	"app/storage/jsondb"
	"fmt"
)

func main() {
	cfg := config.Load()
	strg, err := jsondb.NewConnectionJSON(&cfg)
	if err != nil {
		panic("Failed connect to json:" + err.Error())
	}
	con := controller.NewController(&cfg, strg)
	fmt.Println(con)

	// fmt.Println(con.Task_3("e6ded598-675b-4de2-a1e9-00a876b8e719"))
	fmt.Println(con.Task_9())
}
