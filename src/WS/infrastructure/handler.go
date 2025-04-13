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

func WebSocketHandler(broadcaster domain.Broadcaster) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "No se pudo establecer conexión WebSocket", http.StatusBadRequest)
			return
		}

		// Registrar cliente
		broadcaster.RegisterClient(conn)
		log.Println("👤 Nuevo cliente conectado")

		// Leer mensajes del cliente
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("❌ Error leyendo mensaje:", err)
				broadcaster.UnregisterClient(conn)
				break
			}
			log.Println("📥 Mensaje recibido:", string(message))
			// Enviar el mensaje a todos los clientes conectados
			broadcaster.BroadcastMessage(message)
		}
	}
}

// Función para consumir mensajes de RabbitMQ y transmitir a WebSocket
func ConsumeStockAlertsAndBroadcast(rb domain.RabbitMQ, broadcaster domain.Broadcaster) {
	err := rb.Consume("stock_alerts", "stock.data", func(msg []byte) {
		log.Println("🎧 Mensaje recibido desde RabbitMQ:", string(msg))

		// Aquí puedes realizar alguna lógica adicional si es necesario antes de pasar el mensaje
		// Por ejemplo, verificar si el stock es bajo y solo enviar mensajes de alerta.
		// Si el mensaje contiene una alerta de stock bajo:
		broadcaster.BroadcastMessage(msg) // Enviar el mensaje a todos los clientes conectados
	})

	if err != nil {
		log.Fatalf("❌ Error al consumir mensajes de RabbitMQ: %v", err)
	}
}
