package main

import (
	"fmt"
	"go-chat/handlers"
	"go-chat/room"
	"net/http"
	"os"
)

func main() {
	manager := room.NewManager()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server running on port", port)

	http.HandleFunc("/create", handlers.CreateRoomHandler(manager))
	http.HandleFunc("/ws", handlers.WebSocketHandler(manager))

	http.Handle("/", http.FileServer(http.Dir("./static")))

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Println("Server error:", err)
	}
}
