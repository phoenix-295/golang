package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error
	msg := `Hi, the handshake it complete!`

	for {
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can't send")
		} else {
			fmt.Println("Sending")
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))

	if err := http.ListenAndServeTLS("10.10.10.10:1010", "server.crt", "server.key", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
