package main

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type Message struct {
	Name string      `json:"name"`
	Data interface{} `json:"data"`
}

type Client struct {
	send chan Message
	socket *websocket.Conn
}

func (client *Client) Read {
	var message Message
	for{
		if err := client.socket.ReadJSON(&message); err != nil {
			break
		}
	}

}
func (client *Client) Write() {
	for msg := range client.send {
		fmt.Printf("%#v\n", msg)
		if err := client.socket.WriteJSON(msg); err != {
			break
		}
	}
	_ = client.socket.Close()
}

func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		send:   make(chan Message),
		socket: ws,
	}
}
