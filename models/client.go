package models

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn     *websocket.Conn
	Outgoing chan []byte
	Incoming chan []byte
}

type Message struct {
	Name string
	Body string
}

func (client *Client) Read() {
	message := Message{}
	for {
		err := client.Conn.ReadJSON(&message)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%v\n", message)

		// client.Incoming <- p
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
