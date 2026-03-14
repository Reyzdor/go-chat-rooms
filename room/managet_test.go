package room

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManagerCreateAndGetRoom(t *testing.T) {
	manager := &Manager{
		Rooms: make(map[string]*Room),
	}

	assert.Len(t, manager.Rooms, 0)

	room := manager.CreateRoom()

	assert.Len(t, manager.Rooms, 1)

	fetched := manager.GetRoom(room.Token)
	assert.Equal(t, room, fetched)
	assert.Equal(t, room.Token, fetched.Token)
	assert.Len(t, fetched.Clients, 0)
}

func TestNewManager(t *testing.T) {
	manager := NewManager()
	assert.NotNil(t, manager)
	assert.NotNil(t, manager.Rooms)
	assert.Len(t, manager.Rooms, 0)
}
