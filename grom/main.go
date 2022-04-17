package main

import (
	"fmt"
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	User_id    int
	User_email string
}

type Order_Data struct {
	Order_id   int
	User_email string
	Order_type string
	Price      string
}

var db *gorm.DB

var users []User

var orders []Order_Data

var dsn string = "host=172.19.32.1 user=postgres password=psql dbname=trade_app_2 port=5432 sslmode=disable TimeZone=Asia/Shanghai"

func main() {
	// dsn := "host=172.23.16.1 user=postgres password=postgres dbname=trade_app_2 port=5432 TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	sqlDb, err := db.DB()
	fmt.Println(sqlDb.Stats())
	if err != nil {
		log.Fatalln(err)
	}
	sqlDb.Close()
	// Find all users
	// db.Order("user_id").Find(&users)
	// for _, k := range users {
	// 	fmt.Println(k)
	// }
	// Specific User
	var x string
	fmt.Scanf("%s", &x)
	db.Find(&users, "User_id = ?", x)
	if len(users) == 1 {
		fmt.Println("Got you")
	} else {
		fmt.Println("Email isnt in database")
	}

	//Find all records
	// db.Order("").Find(&users)
	// for _, v := range users {
	// 	fmt.Println(v.User_id, v.User_email)
	// }
	fmt.Println(sqlDb.Stats())
}
