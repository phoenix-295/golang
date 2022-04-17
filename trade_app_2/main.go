package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"
)

//Postgres Data info
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "postgres"
	DB_NAME     = "trade_app"
)

type Order_data struct {
	User_name     string  `json:"u_name"`
	Order_type    string  `json:"order_type"`
	Side          string  `json:"side1"`
	Initial_Size  int     `json:"ins"`
	Size1         int     `json:"size1"`
	Price         float64 `json:"price1"`
	Status        string  `json:"status"`
	Creation_date string  `json:"cr_date"`
	Last_modified string  `json:"lm_date"`
}

var clients = make(map[*websocket.Conn]bool)
var broadcaster = make(chan Order_data)
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

	//load all the messages from DB
	send_buy(ws)

	for {
		var msg Order_data
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

func send_buy(ws *websocket.Conn) {
	db := setup_db()
	// selDB, err := db.Query("SELECT * from order1 where side='Buy'")
	selDB, err := db.Query("SELECT * from order1")
	if err != nil {
		panic(err.Error())
	}
	msg := Order_data{}
	for selDB.Next() {
		var id, user_name, order_type, side, status, creation_date, last_modified string
		var initial_size, size1 int
		var price float64
		err = selDB.Scan(&id, &user_name, &order_type, &side, &initial_size, &size1, &price, &status, &creation_date, &last_modified)
		if err != nil {
			panic(err.Error())
		}
		msg.Side = side
		msg.Size1 = size1
		msg.Price = price
		messageClient(ws, msg)
	}
	defer db.Close()
}

func messageClient(client *websocket.Conn, msg Order_data) {
	err := client.WriteJSON(msg)
	if err != nil && unsafeError(err) {
		log.Printf("error: %v", err)
		client.Close()
		delete(clients, client)
	}
}
func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}

func handleMessages() {
	for {
		// grab any next message from channel
		msg := <-broadcaster
		storeDB(msg)
		messageClients(msg)
	}
}

func storeDB(msg Order_data) {
	db := setup_db()
	defer db.Close()
	currentTime := time.Now()
	insForm, err := db.Prepare(`INSERT INTO order1(user_name, order_type, side, initial_size, size1, price, status, creation_date, last_modified_date) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9)`)
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec("chris_y_k@gmail.com", "limit", msg.Side, 6, msg.Size1, msg.Price, "open", currentTime.Format("2006-01-02 15:04:05"), currentTime.Format("2006-01-02 15:04:05"))
}

func messageClients(msg Order_data) {
	// send to every client currently connected
	for client := range clients {
		messageClient(client, msg)
	}
}

func setup_db() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("public/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})

	r.GET("/ws", func(c *gin.Context) {
		handleConnections(c.Writer, c.Request)
	})
	go handleMessages()

	log.Print("Server starting at http://localhost:9090")

	r.Run()
}
