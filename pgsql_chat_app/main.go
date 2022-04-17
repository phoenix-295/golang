package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "psql"
	DB_NAME     = "chat_app"
)

type Msg_data struct {
	Sender, Message, Timestamp string
}

func setup_db() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	return db
}

func get_rows() []Msg_data {
	db := setup_db()
	rows, err := db.Query("SELECT msg_id, msg_sender, msg_text, timestamp1 FROM chat_logs ORDER BY msg_id DESC")
	checkErr(err)
	// fmt.Println(" Sender | Content | Time ")

	msg_data := Msg_data{}
	msg_datas := []Msg_data{}

	for rows.Next() {
		var msg_id, msg_s, msg_t, t_s string
		err = rows.Scan(&msg_id, &msg_s, &msg_t, &t_s)
		checkErr(err)
		msg_data.Sender = msg_s
		msg_data.Message = msg_t
		msg_data.Timestamp = t_s
		msg_datas = append(msg_datas, msg_data)
		// fmt.Printf("%6v | %6v | %6v \n", msg_s, msg_t, t_s)
	}
	defer db.Close()
	return msg_datas
}

// func indexHandler(w http.ResponseWriter, r *http.Request) {
// 	var tmpl = template.Must(template.ParseFiles("login.html"))
// 	// t, _ := template.ParseFiles("login.html")
// 	tmpl.ExecuteTemplate(w, "Index", nil)
// }

func send(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("send_message.html"))
	table := get_rows()
	tmpl.ExecuteTemplate(w, "Index", table)
}

func view(w http.ResponseWriter, r *http.Request) {
	var tmpl = template.Must(template.ParseFiles("view_mess.html"))
	table := get_rows()
	tmpl.ExecuteTemplate(w, "Index", table)
}

// func main() {
// 	fmt.Println("Server at 8080")
// 	http.HandleFunc("/", indexHandler)
// 	// http.HandleFunc("/send", send)
// 	// http.HandleFunc("/view", view)
// 	http.ListenAndServe(":3000", nil)

// }

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
