# Go-Chat 🚀

A simple messenger written in Go that allows real-time communication through private rooms. It works on the principle of creating temporary chat rooms with a unique token. The first user creates a room and receives a token, then shares it with the second user through any channel. Both users connect to the server using this token and chosen nicknames, and once connected they can exchange messages that are instantly delivered to all room participants.

The core is built on WebSocket connections providing persistent two-way communication between client and server without the need for constant reconnections. The server is written in Go using the gorilla/websocket library and stores all active rooms in memory. The frontend is built with pure HTML, CSS and JavaScript without any third-party frameworks.


![image alt](https://github.com/Reyzdor/go-chat-rooms/blob/632f758e5dd84acece453af15af8b3d29ede6996/1.jpg)


## WebSocket Connection

In the WebSocketHandler function, the connection is created by upgrading the HTTP request:

```go
conn, err := upgrader.Upgrade(w, r, nil)
if err != nil {
    return
}
```

This connection then persists until the client disconnects. All the magic happens in the upgrader, which turns regular HTTP into a persistent WebSocket.

## Room Creation and Token System

When CreateRoomHandler is called:

```go
room := manager.CreateRoom()
w.Header().Set("Content-Type", "application/json")
json.NewEncoder(w).Encode(CreateRoomResponse{
    Token: room.Token,
})
```

And inside manager.CreateRoom() a token is generated and the room is saved in the manager's map:

```go
token := generateToken()
room := &Room{
    Token:   token,
    Clients: make(map[*Client]bool),
}
m.Rooms[token] = room
```

## Adding Users to a Room

In joinRoom, a WebSocket with a token and nickname is created on the client:

```js
ws = new WebSocket(`ws://localhost:8080/ws?nick=${nick}&token=${token}`);
```

On the server, WebSocketHandler catches this:

```go
token := r.URL.Query().Get("token")
nick := r.URL.Query().Get("nick")
rm := manager.GetRoom(token)
client := &room.Client{
    Nick: nick,
    Conn: conn,
}
rm.Clients[client] = true
```

The client is added to the room map. The ```map[*Client]bool``` map is used as an array to quickly add and remove participants.

## Message Broadcasting

When a message arrives from a client:

```go
_, msg, err := conn.ReadMessage()
fullMessage := client.Nick + ": " + string(msg)
for c := range rm.Clients {
    c.Conn.WriteMessage(websocket.TextMessage, []byte(fullMessage))
}
```

The server runs through all clients in the room map and sends a message to each one. This is detected on the client via ```ws.onmessage```.

## Connection Cleanup

```go
defer func() {
    delete(rm.Clients, client)
    conn.Close()
}()
```

## Quick Start

```bash
# Clone repository
git clone https://github.com/Reyzdor/go-chat-rooms.git

cd go-chat-rooms
```

If for some reason an error occurs with ```go.mod``` and ```go.sum```, simply remove them from your project structure and run the commands below:

```bash
# Initialize the module container
go mod init go-chat

# Install library
go get github.com/gorilla/websocket
```

Start the server:

```bash
go run main.go
```

Open browser: ```http://localhost:8080```

The go.mod and go.sum files are kept in the repository as they are required for proper dependency management and ensure the project builds consistently for all developers.
