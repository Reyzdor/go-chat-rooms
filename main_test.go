package main

import (
	"go-chat/handlers"
	"go-chat/room"
	"net/http"
	"testing"
)

func TestMainHandlers(t *testing.T) {
	manager := room.NewManager()
	_ = handlers.CreateRoomHandler(manager)
	_ = handlers.WebSocketHandler(manager)
	http.DefaultServeMux = http.NewServeMux()
}
