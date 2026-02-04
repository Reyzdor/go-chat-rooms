package room

import (
	"crypto/rand"
	"encoding/hex"
	"sync"
)

type Manager struct {
	Rooms map[string]*Room
	Mutex sync.Mutex
}

func NewManager() *Manager {
	return &Manager{
		Rooms: make(map[string]*Room),
	}
}

func generateToken() string {
	b := make([]byte, 4)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (m *Manager) CreateRoom() *Room {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	token := generateToken()

	room := &Room{
		Token:   token,
		Clients: make(map[*Client]bool),
	}

	m.Rooms[token] = room
	return room
}

func (m *Manager) GetRoom(token string) *Room {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	return m.Rooms[token]
}
