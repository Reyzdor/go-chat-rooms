package room

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Client struct {
	Nick string
	Conn *websocket.Conn
}

type Room struct {
	Token   string
	Clients map[*Client]bool
	Mutex   sync.Mutex
}
