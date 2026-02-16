package handlers

import (
	"go-chat/go-chat/room"
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
			http.Error(w, "Комната не найдена", http.StatusNotFound)
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

		rm.Mutex.Lock()
		rm.Clients[client] = true
		rm.Mutex.Unlock()

		defer func() {
			rm.Mutex.Lock()
			delete(rm.Clients, client)
			rm.Mutex.Unlock()
			conn.Close()
		}()

		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}

			fullMessage := client.Nick + ": " + string(msg)

			rm.Mutex.Lock()

			clients := make([]*room.Client, 0, len(rm.Clients))
			for c := range rm.Clients {
				clients = append(clients, c)
			}
			rm.Mutex.Unlock()

			for _, c := range clients {
				c.Conn.WriteMessage(websocket.TextMessage, []byte(fullMessage))
			}
		}
	}
}
