package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

var login1 = false
var g_u_name string

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "chat_app"
)

type Msg_data struct {
	Sender     string `json:"msg_sender"`
	Message    string `json:"msg_text"`
	Timestamp1 string `json:"timestamp1"`
}

//new code
var clients = make(map[*websocket.Conn]bool)
var broadcaster = make(chan Msg_data)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// ensure connection close when function returns
	defer ws.Close()
	clients[ws] = true

	// if it's zero, no messages were ever sent/saved
	// if rdb.Exists("chat_messages").Val() != 0 {
	// 	sendPreviousMessages(ws)
	// }

	for {
		var msg Msg_data
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			delete(clients, ws)
			break
		}
		// send new message to the channel
		broadcaster <- msg
	}
}

func sendPreviousMessages(ws *websocket.Conn) {
	chatMessages, err := rdb.LRange("chat_messages", 0, -1).Result()
	if err != nil {
		panic(err)
	}

	// send previous messages
	for _, chatMessage := range chatMessages {
		var msg Msg_data
		json.Unmarshal([]byte(chatMessage), &msg)
		messageClient(ws, msg)
		fmt.Println(msg)
	}
}

// // If a message is sent while a client is closing, ignore the error
// func unsafeError(err error) bool {
// 	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
// }

// func handleMessages() {
// 	for {
// 		// grab any next message from channel
// 		msg := <-broadcaster

// 		storeInRedis(msg)
// 		messageClients(msg)
// 	}
// }

// //pg database connection
// // func setup_db() *sql.DB {
// // 	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
// // 		DB_USER, DB_PASSWORD, DB_NAME)
// // 	db, err := sql.Open("postgres", dbinfo)
// // 	if err != nil {
// // 		panic(err.Error())
// // 	}
// // 	return db
// // }

// func storeInRedis(msg Msg_data) {
// 	json, err := json.Marshal(msg)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if err := rdb.RPush("chat_messages", json).Err(); err != nil {
// 		panic(err)
// 	}
// }

// func messageClients(msg Msg_data) {
// 	// send to every client currently connected
// 	for client := range clients {
// 		messageClient(client, msg)
// 	}
// }

// func messageClient(client *websocket.Conn, msg Msg_data) {
// 	err := client.WriteJSON(msg)
// 	if err != nil && unsafeError(err) {
// 		log.Printf("error: %v", err)
// 		client.Close()
// 		delete(clients, client)
// 	}
// }

//end

func setup_db() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err.Error())
	}
	return db
}

var tmpl = template.Must(template.ParseGlob("form/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	// conn, _, _, err := ws.UpgradeHTTP(r, w)
	tmpl.ExecuteTemplate(w, "Index", nil)
}

func get_data() {
	db := setup_db()
	selDB, err := db.Query("SELECT * FROM chat_logs ORDER BY msg_id DESC")
	if err != nil {
		panic(err.Error())
	}
	msg := Msg_data{}
	msgs := []Msg_data{}
	for selDB.Next() {
		var id, sender, message, timestamp1 string
		err = selDB.Scan(&id, &sender, &message, &timestamp1)
		if err != nil {
			panic(err.Error())
		}
		msg.Sender = sender
		msg.Message = message
		msg.Timestamp1 = timestamp1
		msgs = append(msgs, msg)
	}
	defer db.Close()
}

func Show(w http.ResponseWriter, r *http.Request) {
	if login1 == true {
		db := setup_db()
		selDB, err := db.Query("SELECT * FROM chat_logs ORDER BY msg_id DESC")
		if err != nil {
			panic(err.Error())
		}
		msg := Msg_data{}
		msgs := []Msg_data{}
		for selDB.Next() {
			var id, sender, message, timestamp1 string
			err = selDB.Scan(&id, &sender, &message, &timestamp1)
			if err != nil {
				panic(err.Error())
			}
			msg.Sender = sender
			msg.Message = message
			msg.Timestamp1 = timestamp1
			msgs = append(msgs, msg)
		}
		tmpl.ExecuteTemplate(w, "Show", msgs)
		defer db.Close()
	} else {
		http.Redirect(w, r, "/", 301)
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	var timestamp1 = time.Now()
	db := setup_db()
	if r.Method == "POST" {

		message := r.FormValue("message")
		insForm, err := db.Prepare(`INSERT INTO chat_logs(msg_sender, msg_text, timestamp1) VALUES($1,$2,$3)`)
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(g_u_name, message, timestamp1)
	}
	fmt.Println("Inserted 1 record successfully")
	defer db.Close()
	http.Redirect(w, r, "/show", 301)
}

func Login(w http.ResponseWriter, r *http.Request) {
	db := setup_db()
	var y string
	if r.Method == "POST" {
		u_name := r.FormValue("uname")
		pass := r.FormValue("psw")
		g_u_name = r.FormValue("uname")
		fmt.Println("Username: ", u_name)
		fmt.Println("Password: ", pass)
		err := db.QueryRow("Select * From chat_users where user_name=$1 AND user_pass=$2", u_name, pass).Scan(&y)
		if err != nil && err != sql.ErrNoRows {
			fmt.Println("Hit")
			login1 = true
			http.Redirect(w, r, "/show", 301)
		} else {
			fmt.Println("Fail")
			login1 = false
			http.Redirect(w, r, "/", 301)
		}
	}
	defer db.Close()
}

func main() {
	fmt.Println("Started at port http://localhost:1234/")
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/login", Login)
	http.ListenAndServe(":1234", nil)
}
