package main

import (
	"fmt"
	"net"
	"time"
)

var connections = make([]net.Conn, 0)

func main() {
	listener, _ := net.Listen("tcp", ":6666")
	go sendMessage()

	for {
		newConnection, _ := listener.Accept()

		connections = append(connections, newConnection)
	}
}

func sendMessage() {
	timer := time.Tick(2 * time.Second)

	for c := range timer {
		toSend := fmt.Sprintf("%s\n", c.String())

		for _, connection := range connections {
			connection.Write([]byte(toSend))
		}
	}
}
