package handlers

import (
	"encoding/json"
	"go-chat/room"
	"net/http"
)

type CreateRoomResponse struct {
	Token string `json:"token"`
}

func CreateRoomHandler(manager *room.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		room := manager.CreaetRoom()

		json.NewEncoder(w).Encode(CreateRoomResponse{
			Token: room.Token,
		})
	}
}
