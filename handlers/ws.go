package handlers

import (
	"go-chat/room"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func WebSocketHandler(manager *room.Manager) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		nick := r.URL.Query().Get("nick")

		rm := manager.GetRoom(token)
		if rm == nil {
			http.Error(w, "Комната не найдена!", http.StatusNotFound)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		client := &room.Client{
			Nick: nick,
			Conn: conn,
		}

		rm.Clients[client] = true

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				delete(rm.Clients, client)
				return
			}

			fullMessage := client.Nick + ": " + string(msg)

			for c := range rm.Clients {
				c.Conn.WriteMessage(
					websocket.TextMessage,
					[]byte(fullMessage),
				)
			}
		}
	}
}
