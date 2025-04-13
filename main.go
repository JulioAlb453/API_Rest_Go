package main

import (
	ws "API_ejemplo/src/WS"
	"API_ejemplo/src/WS/infrastructure"
	"API_ejemplo/src/routes"
	"log"
	"net/http"
)

func main() {
	broadcaster := ws.NewWebSocketBroadcaster()

	go func() {
		mux := http.NewServeMux()
		mux.HandleFunc("/ws", infrastructure.WebSocketHandler(broadcaster))

		log.Println("🧠 WebSocket escuchando en :8081/ws")
		err := http.ListenAndServe(":8081", mux)
		if err != nil {
			log.Fatalf("❌ Error iniciando WebSocket server: %v", err)
		}
	}()

	// Gin REST API en :8080
	router := routes.SetupRouter()
	log.Println("🚀 API REST escuchando en :8080")
	router.Run(":8080")
}
