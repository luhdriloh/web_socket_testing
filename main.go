package main

import (
	"github.com/gorilla/websocket"
	"github.com/luhdriloh/http/models"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var chatroom = models.NewChatroom()

func main() {
	http.HandleFunc("/socket", socketHandler)
	http.ListenAndServe(":8000", nil)
}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	addClient(conn)
}

func addClient(conn *websocket.Conn) {
	chatroom.NewClient <- conn
}
