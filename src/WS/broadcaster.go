// src/album/infrastructure/ws/broadcaster.go

package ws

import (
	"API_ejemplo/src/album/domain"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketBroadcaster struct {
	clients map[string]*websocket.Conn  // Usamos un mapa con UserID como clave
	mu      sync.Mutex
}

var _ domain.Broadcaster = &WebSocketBroadcaster{}

// NewWebSocketBroadcaster crea una nueva instancia de WebSocketBroadcaster
func NewWebSocketBroadcaster() *WebSocketBroadcaster {
	return &WebSocketBroadcaster{
		clients: make(map[string]*websocket.Conn),  // Cambiar la estructura para usar UserID
	}
}

// RegisterClient registra un cliente por UserID y conexión
func (b *WebSocketBroadcaster) RegisterClient(client domain.Client) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.clients[client.UserID] = client.Connection
	log.Println("👤 Cliente registrado:", client.UserID)
}

// UnregisterClient desregistra un cliente por UserID
func (b *WebSocketBroadcaster) UnregisterClient(client domain.Client) {
	b.mu.Lock()
	defer b.mu.Unlock()
	client.Connection.Close()
	delete(b.clients, client.UserID)
	log.Println("👤 Cliente desregistrado:", client.UserID)
}

// BroadcastMessage envía un mensaje a todos los clientes conectados
func (b *WebSocketBroadcaster) BroadcastMessage(message []byte) {
	b.mu.Lock()
	defer b.mu.Unlock()

	log.Println("🔄 Broadcasting message to all clients:", string(message)) // Log antes de enviar el mensaje

	for userID, client := range b.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("❌ Error enviando mensaje a cliente:", userID, err)
			client.Close()
			delete(b.clients, userID)
		} else {
			log.Println("📤 Mensaje enviado a cliente:", userID) // Log cuando se envía el mensaje a cada cliente
		}
	}
}
