package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/socket", socketHandler)
	http.ListenAndServe(":8000", nil)
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	go sendMessage(conn)
}

func sendMessage(conn *websocket.Conn) {
	timer := time.Tick(2 * time.Second)
	messageType := websocket.TextMessage

	for c := range timer {
		toSend := []byte(fmt.Sprintf("%s\n", c.String()))

		if err := conn.WriteMessage(messageType, toSend); err != nil {
			return
		}
	}
}
