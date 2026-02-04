package main

import (
	"go-chat/handlers"
	"go-chat/room"
	"net/http"
)

func main() {
	manager := room.NewManager()

	http.HandleFunc("/create", handlers.CreateRoomHandler(manager))
	http.HandleFunc("/ws", handlers.WebSocketHandler(manager))

	http.Handle("/", http.FileServer(http.Dir("./static")))

	http.ListenAndServe(":8080", nil)
}
