package handlers

import (
	"go-chat/room"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
)

func TestWebSocketHandlerRoomNotFound(t *testing.T) {
	manager := room.NewManager()
	req := httptest.NewRequest("GET", "/ws?token=noexistent&nick=Roma", nil)
	w := httptest.NewRecorder()

	handler := WebSocketHandler(manager)
	handler(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusNotFound, resp.StatusCode)
}

func TestWebSocketHandlerSuccess(t *testing.T) {
	manager := room.NewManager()
	rm := manager.CreateRoom()

	srv := httptest.NewServer(WebSocketHandler(manager))
	defer srv.Close()

	wsURL := "ws" + srv.URL[len("http"):] + "?token=" + rm.Token + "&nick=Tester"

	conn, resp, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.NoError(t, err)
	defer conn.Close()
	assert.Equal(t, http.StatusSwitchingProtocols, resp.StatusCode)

	time.Sleep(50 * time.Millisecond)
	found := false
	rm.Mutex.Lock()
	for c := range rm.Clients {
		if c.Nick == "Tester" {
			found = true
		}
	}
	rm.Mutex.Unlock()
	assert.True(t, found)
}

func TestWebSocketHandlerSendReceive(t *testing.T) {
	manager := room.NewManager()
	rm := manager.CreateRoom()

	srv := httptest.NewServer(WebSocketHandler(manager))
	defer srv.Close()

	wsURL := "ws" + srv.URL[len("http"):] + "?token=" + rm.Token + "&nick=Tester"

	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	assert.NoError(t, err)
	defer conn.Close()

	err = conn.WriteMessage(websocket.TextMessage, []byte("Hello"))
	assert.NoError(t, err)

	conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
	_, msg, err := conn.ReadMessage()
	assert.NoError(t, err)
	assert.Contains(t, string(msg), "Tester: Hello")
}

func TestWebSocketHandlerUpgradeError(t *testing.T) {
	manager := room.NewManager()
	manager.CreateRoom()

	req := httptest.NewRequest("GET", "/ws?token=invalid&nick=test", nil)
	w := httptest.NewRecorder()

	handler := WebSocketHandler(manager)
	handler(w, req)
}
