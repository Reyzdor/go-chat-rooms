package handlers

import (
	"encoding/json"
	"go-chat/room"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateRoomHandler(t *testing.T) {
	manager := room.NewManager()

	req := httptest.NewRequest("GET", "/create", nil)
	w := httptest.NewRecorder()

	handler := CreateRoomHandler(manager)
	handler(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

	var body CreateRoomResponse
	err := json.NewDecoder(resp.Body).Decode(&body)
	assert.NoError(t, err)
	assert.NotEmpty(t, body.Token)

	room := manager.GetRoom(body.Token)
	assert.NotNil(t, room)
	assert.Equal(t, body.Token, room.Token)
}
