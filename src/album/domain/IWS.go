// src/album/domain/broadcaster.go

package domain

import "github.com/gorilla/websocket"

// Broadcaster define las operaciones para gestionar clientes WebSocket y la difusión de mensajes
type Broadcaster interface {
	RegisterClient(client Client)
	UnregisterClient(client Client)
	BroadcastMessage(msg []byte)
}

// Client representa un cliente conectado, con su UserID y la conexión WebSocket
type Client struct {
	UserID     string
	Connection *websocket.Conn
}
