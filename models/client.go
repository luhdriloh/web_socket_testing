package models

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Outgoing chan []byte
	Incoming chan []byte
}

func (client *Client) Read() {
	for {
		_, p, err := client.Conn.ReadMessage()
		if err != nil {
			return
		}

		client.Incoming <- p
	}
}

func (client *Client) Write(toSend []byte) {
	messageType := websocket.TextMessage

	err := client.Conn.WriteMessage(messageType, toSend)
	if err != nil {
		return
	}
}

func (client *Client) Listen() {
	go client.Read()
}

func NewClient(connection *websocket.Conn) *Client {
	client := &Client{
		Conn:     connection,
		Outgoing: make(chan []byte),
		Incoming: make(chan []byte),
	}

	client.Listen()

	return client
}
