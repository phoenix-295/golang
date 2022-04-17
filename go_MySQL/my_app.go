package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/test1")
	if err != nil {
		panic(err.Error())
	}
	rows, err := db.Query("Select * from food_category")

	checkError(err)

	for rows.Next() {
		var cat_id int
		var cat_name string
		var cat_details string
		var space1 string = "\t"

		err = rows.Scan(&cat_id, &cat_name, &cat_details)
		checkError(err)
		fmt.Print(cat_id, space1)
		fmt.Print(cat_name, space1)
		fmt.Print(cat_details, space1)
		fmt.Println()
	}
	defer db.Close()
}
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
