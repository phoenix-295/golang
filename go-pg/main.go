package main

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

type Users struct {
	User_id    int
	User_email string
}

type Order_data struct {
	O_type             string  `json:"my_order"`
	Order_id           int     `json:"oid1"`
	User_email         string  `json:"u_name"`
	Order_type         string  `json:"order_type"`
	Side               string  `json:"side1"`
	Initial_size       float64 `json:"ins"`
	Size               float64 `json:"size1"`
	Price              float64 `json:"price1"`
	Status             string  `json:"status"`
	Creation_date      string  `json:"cr_date"`
	Last_modified_date string  `json:"lm_date"`
}

func main() {
	db := pg.Connect(&pg.Options{
		Addr:     "winhost:5432",
		User:     "postgres",
		Password: "psql",
		Database: "trade_app_2",
	})
	defer db.Close()

	// SELECT * FROM users WHERE "user_id" = '3'
	users := new(Users)
	err := db.Model(users).ColumnExpr("*").Where("? = ?", pg.Ident("user_id"), "3").Select()
	chkErr(err)
	fmt.Println(users)

	// SELECT * FROM order_data WHERE user_email = 'eric_k04@gmail.com'
	// u_email := users.User_email
	// orders := new([]Order_data)
	// err := db.Model(orders).ColumnExpr("*").Where("? = ?", pg.Ident("user_email"), u_email).Select()
	// chkErr(err)
	// for _, v := range *orders {
	// 	fmt.Println(v.Order_id)
	// }
	// fmt.Println(len(*orders))
}

func chkErr(e error) {
	if e != nil {
		panic(e)
	}
}
