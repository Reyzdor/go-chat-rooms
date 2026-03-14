package room

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClient(t *testing.T) {
	client := Client{Nick: "Nick"}

	assert.Equal(t, "Nick", client.Nick)

	assert.NotEmpty(t, client.Nick)

	assert.Nil(t, client.Conn)
}

func TestRoomCreate(t *testing.T) {
	room := Room{
		Token:   "123room",
		Clients: make(map[*Client]bool),
	}

	assert.Equal(t, "123room", room.Token)
	assert.Len(t, room.Clients, 0)
}

func TestRoomAddClient(t *testing.T) {
	room := Room{
		Token:   "room1",
		Clients: make(map[*Client]bool),
	}

	client := &Client{
		Nick: "Roman",
	}

	room.Mutex.Lock()
	room.Clients[client] = true
	room.Mutex.Unlock()

	assert.Len(t, room.Clients, 1)
	assert.True(t, room.Clients[client])
}
