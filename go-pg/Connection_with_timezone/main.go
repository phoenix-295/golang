package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
)

type Demo struct {
	tableName  struct{} `pg:"test_data"`
	Id         int
	Time_stamp string
}

var db *pg.DB

func setup_pg() *pg.DB {
	err := godotenv.Load()
	checkError(err)
	DB_USER := os.Getenv("DB_USER")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	db := pg.Connect(&pg.Options{
		Addr:     DB_HOST,
		User:     DB_USER,
		Password: DB_PASSWORD,
		Database: DB_NAME,
	})

	checkError(err)
	return db
}

func main() {
	db = setup_pg()
	defer db.Close()
	var err error

	// db.Options().OnConnect(ctx, db.Conn()) //*pg.DB

	// _, err = db.Model().Exec(
	// 	"SET timezone = 'UTC';",
	// )
	// checkError(err)

	var ins Demo
	ins.Time_stamp = epochToHumanReadable(time.Now().UnixMilli(), 0).String()[:29]
	fmt.Println(time.UnixMilli(time.Now().UnixMilli()).String()[:29])
	_, err = db.Model(&ins).Insert()
	checkError(err)

	var tp []Demo
	db.Model().Exec("SET timezone = 'UTC'")
	checkError(err)
	db.Model(&tp).Select()
	fmt.Println("tp", tp[len(tp)-1].Time_stamp)
}

func checkError(err error) (b bool) {
	if err != nil {
		_, fn, line, _ := runtime.Caller(1)
		log.Printf("[error] %s:%d %v", fn, line, err)
		b = true
	}
	return
}

func epochToHumanReadable(epoch int64, epoch2 int64) time.Time {
	return time.UnixMilli(epoch)
}
