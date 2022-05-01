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

type User_accounts struct {
	Acc_id                  int
	User_email              string
	Mobile_no               int
	Creation_date           string
	Last_modified_date      string
	Wallet_address          string
	BTC_balance             float64 `pg:"BTC_balance"`
	INR_balance             float64 `pg:"INR_balance" json:"lm_date"`
	Collateral              float64
	Locked_BTC_balance      float64
	Locked_INR_balance      float64
	Unconfirmed_BTC_balance float64
	Payment_code            string
	Fee_discount            string
	Kyc_level               int
	Otp_secret              string
}

// Django Trade
type Order_entry_user struct {
	tableName    struct{} `pg:"order_entry_user"`
	Id           int
	Password     string
	Last_login   string
	Is_superuser bool
	Username     string
	First_name   string
	Last_name    string
	Email        string
	Is_staff     bool
	Is_active    bool
	Date_joined  string
}

type Order_entry_trade struct {
	tableName          struct{} `pg:"order_entry_trade"`
	Id                 int
	Trade_size         float64
	Trade_price        float64
	Is_uptrade         bool
	Buyer_fee          float64
	Seller_fee         float64
	Creation_date      string
	Last_modified_date string
	Buy_order_id       int
	Sell_order_id      int
}

type Order_entry_useraccount struct {
	tableName               struct{} `pg:"order_entry_useraccount"`
	User_id                 int
	Mobile_no               int
	Creation_date           string
	Last_modified_date      string
	Wallet_address          string
	BTC_balance             float64 `pg:"BTC_balance"`
	INR_balance             float64 `json:"lm_date" pg:"INR_balance"`
	Collateral              float64
	Locked_BTC_balance      float64 `pg:"locked_BTC_balance"`
	Locked_INR_balance      float64 `pg:"locked_INR_balance"`
	Unconfirmed_BTC_balance float64 `pg:"unconfirmed_BTC_balance"`
	Payment_code            string
	Fee_discount            float64
	Kyc_level               int16
	Otp_secret              string
}

type Order_entry_order struct {
	tableName          struct{} `pg:"order_entry_order"`
	Id                 int
	Side               string
	Size               float64
	Price              float64
	Creation_date      string
	Last_modified_date string
	User_id            string
	Initial_size       float64
	Status             string
	Order_type         string
}

type Data struct {
	Id   int
	Name string
}

func main() {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "django_trade",
	})
	defer db.Close()

	// // SELECT * FROM users WHERE "user_id" = '3'
	// users := new(Users)
	// err := db.Model(users).ColumnExpr("*").Where("? = ?", pg.Ident("user_id"), "3").Select()
	// chkErr(err)
	// fmt.Println(users.User_email)

	// // SELECT * FROM order_data WHERE user_email = 'eric_k04@gmail.com'
	// u_email := users.User_email
	// orders := new([]Order_data)
	// err2 := db.Model(orders).ColumnExpr("*").Where("? = ?", pg.Ident("user_email"), u_email).Select()
	// chkErr(err2)
	// for _, v := range *orders {
	// 	fmt.Println(v.Order_id)
	// }
	// fmt.Println(len(*orders))

	// // SELECT * from order_data WHERE user_email = $1 and status='Open' and side='Sell' ORDER BY price limit 20`
	// od := new([]Order_data)
	// err := db.Model(od).Where("user_email=?", "tom2121@gmail.com").Where("Status='Open'").Where("Side='Sell'").Order("price").Limit(20).Select()
	// chkErr(err)
	// for _, v := range *od {
	// 	fmt.Println("Order", v)
	// }

	// // SELECT acc_id,locked_btc_balance,locked_inr_balance FROM user_accounts WHERE user_email= $1
	// user_accounts := new(User_accounts)
	// err := db.Model(user_accounts).Where("user_email='tom2121@gmail.com'").Select()
	// chkErr(err)
	// fmt.Println("Account", user_accounts)
	// fmt.Println("Before substraction", user_accounts.Locked_BTC_balance)

	// user_accounts.Locked_BTC_balance -= 2
	// fmt.Println("After substraction", user_accounts.Locked_BTC_balance)
	// log.Printf("Hello")
	// // Update user_accounts set locked_btc_balance=$1 WHERE acc_id=$2`, acc_id, locked_BTC_balance)
	// _, err2 := db.Model(user_accounts).Set("locked_btc_balance=?", user_accounts.Locked_BTC_balance).Where("acc_id=?", user_accounts.Acc_id).Update()
	// chkErr(err2)

	// // // Django
	// user := new([]Order_entry_user)
	// err := db.Model(user).Select()
	// chkErr(err)
	// for _, v := range *user {
	// 	fmt.Println("", v)
	// }

	// trade := new([]Order_entry_trade)
	// err1 := db.Model(trade).Select()
	// chkErr(err1)
	// for _, v := range *trade {
	// 	fmt.Println("", v)
	// }

	// user_account := new([]Order_entry_useraccount)
	// err2 := db.Model(user_account).Select()
	// chkErr(err2)
	// for _, v := range *user_account {
	// 	fmt.Println("", v)
	// }

	// order_entry := new([]Order_entry_order)
	// err3 := db.Model(order_entry).Where("side='Sell'").Where("status='Open'").Where("user_id!=?", 15).Order("price").Order("last_modified_date").Select()
	// chkErr(err3)
	// for _, v := range *order_entry {
	// 	fmt.Println("", v)
	// }

	var d Data
	d.Name = "Nikhil"
	var res int
	_, err4 := db.Model(&d).Returning("id").Insert(&res)
	chkErr(err4)
	fmt.Println("X", res)
}

func chkErr(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
