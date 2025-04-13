package domain

import "github.com/gorilla/websocket"

type Broadcaster interface {
	RegisterClient(conn *websocket.Conn)
	UnregisterClient(conn *websocket.Conn)
	BroadcastMessage(msg []byte)
}
