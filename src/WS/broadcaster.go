package ws

import (
	"API_ejemplo/src/album/domain"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type WebSocketBroadcaster struct {
	clients map[*websocket.Conn]bool
	mu      sync.Mutex
}

var _ domain.Broadcaster = &WebSocketBroadcaster{}

func NewWebSocketBroadcaster() *WebSocketBroadcaster {
	return &WebSocketBroadcaster{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (b *WebSocketBroadcaster) RegisterClient(client *websocket.Conn) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.clients[client] = true
}

func (b *WebSocketBroadcaster) UnregisterClient(client *websocket.Conn) {
	b.mu.Lock()
	defer b.mu.Unlock()
	client.Close()
	delete(b.clients, client)
}

func (b *WebSocketBroadcaster) BroadcastMessage(message []byte) {
	b.mu.Lock()
	defer b.mu.Unlock()
	
	log.Println("🔄 Broadcasting message to all clients:", string(message)) // Log antes de enviar el mensaje
	
	for client := range b.clients {
		err := client.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Println("❌ Error enviando mensaje a cliente:", err)
			client.Close()
			delete(b.clients, client)
		} else {
			log.Println("📤 Mensaje enviado a cliente:", client.RemoteAddr()) // Log cuando se envía el mensaje a cada cliente
		}
	}
}
