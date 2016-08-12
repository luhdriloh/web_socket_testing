package models

import (
	"github.com/gorilla/websocket"
)

type Chatroom struct {
	Clients   []*Client
	Incoming  chan []byte
	NewClient chan *websocket.Conn
}

// we need to listen to all incoming connection and add to client
// listen to any incoming messages and send them to all the clients

func (chatroom *Chatroom) Broadcast(data []byte) {
	for _, client := range chatroom.Clients {
		client.Write(data)
	}
}

func (chatroom *Chatroom) Join(connection *websocket.Conn) {
	client := NewClient(connection)
	chatroom.Clients = append(chatroom.Clients, client)

	// listen to any incoming client messages
	go func() {
		for {
			chatroom.Incoming <- <-client.Incoming
		}
	}()
}

func (chatroom *Chatroom) Listen() {
	go func() {
		for {
			select {
			case data := <-chatroom.Incoming:
				chatroom.Broadcast(data)
			case connection := <-chatroom.NewClient:
				chatroom.Join(connection)
			}
		}
	}()
}

func NewChatroom() *Chatroom {
	chatroom := &Chatroom{
		Clients:   make([]*Client, 0),
		Incoming:  make(chan []byte),
		NewClient: make(chan *websocket.Conn),
	}

	chatroom.Listen()

	return chatroom
}
