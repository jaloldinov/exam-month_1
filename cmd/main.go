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

	// con.Task_4()
	con.Task_5()

}
