package room

import "github.com/gorilla/websocket"

type Client struct {
	Nick string
	Conn *websocket.Conn
}

type Room struct {
	Token   string
	Clients map[*Client]bool
}
