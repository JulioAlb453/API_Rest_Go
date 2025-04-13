// src/album/infrastructure/ws/handler.go

package infrastructure

import (
	"API_ejemplo/src/album/domain"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },  // Permitir conexiones desde cualquier origen
}

// WebSocketHandler maneja las conexiones WebSocket y los mensajes de los clientes
func WebSocketHandler(broadcaster domain.Broadcaster) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.URL.Query().Get("userID")
		if userID == "" {
			http.Error(w, "Se requiere un UserID", http.StatusBadRequest)
			return
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "No se pudo establecer conexión WebSocket", http.StatusBadRequest)
			return
		}

		// Registrar cliente
		client := domain.Client{
			UserID:     userID,
			Connection: conn,
		}
		broadcaster.RegisterClient(client)
		log.Println("👤 Nuevo cliente conectado:", userID)

		// Leer mensajes del cliente
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("❌ Error leyendo mensaje de cliente:", userID, err)
				broadcaster.UnregisterClient(client)
				break
			}
			log.Println("📥 Mensaje recibido de cliente:", userID, string(message))

			// Enviar el mensaje a todos los clientes conectados
			broadcaster.BroadcastMessage(message)
		}
	}
}
