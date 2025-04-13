package infraestructure

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	ws "API_ejemplo/src/WS"
	"API_ejemplo/src/WS/infrastructure"
	"API_ejemplo/src/album/application"
	"API_ejemplo/src/album/domain"
	"API_ejemplo/src/album/infraestructure/controllers"
	"API_ejemplo/src/album/infraestructure/repository"
	"API_ejemplo/src/core"
	"API_ejemplo/src/shared/broker"
)

type Dependencies struct {
	AlbumSaveController        *controllers.AlbumSaveController
	AlbumGetByIdController     *controllers.AlbumGetByIdController
	AlbumGetByArtistController *controllers.AlbumGetByArtistController
	AlbumGetByTitleController  *controllers.AlbumGetByTitleController
	AlbumGetAllController      *controllers.AlbumGetAllController
	AlbumUpdateController      *controllers.AlbumUpdateController
	AlbumDeleteController      *controllers.AlbumDeleteController

	ShortPollingStockController *controllers.ShortPollingStockController
	LongPollingController       *controllers.LongPollingController
	ShortPollingPriceController *controllers.ShortPollingPriceController

	Broadcaster *ws.WebSocketBroadcaster
}

func Init() *Dependencies {
	conn := core.Connect()
	if conn == nil {
		log.Fatal("❌ Error al conectar con la base de datos")
	}
	db := conn.Database("MundyWalk")
	client := conn

	rb, err := broker.NewRabbitMQBroker("amqp://guest:guest@3.209.113.62:5672/")
	if err != nil {
		log.Fatalf("❌ Fallo al conectar a RabbitMQ: %v", err)
	}

	albumRepo := repository.NewMongoAlbumRepository(db)

	broadcaster := ws.NewWebSocketBroadcaster()

	http.HandleFunc("/ws", infrastructure.WebSocketHandler(broadcaster))

	updateAlbumUseCase := application.NewUpdateAlbumsUseCase(albumRepo, rb, broadcaster)

	err = rb.Consume("stock_alerts", "stock.data", func(msg []byte) {
		log.Println("🎧 Mensaje recibido desde RabbitMQ:", string(msg))
	
		var album domain.Album
		if err := json.Unmarshal(msg, &album); err != nil {
			log.Printf("❌ Error deserializando mensaje: %v", err)
			return
		}
	
		if _, err := updateAlbumUseCase.Execute(context.Background(), album); err != nil {
			log.Printf("⚠️ Error procesando mensaje: %v", err)
		}
	
		broadcaster.BroadcastMessage(msg)
	})
	

	if err != nil {
		log.Fatalf("❌ Error al consumir mensajes de RabbitMQ: %v", err)
	}

	createAlbumUseCase := application.NewCreatedAlbumUseCase(albumRepo)
	getAlbumByIdUseCase := application.NewGetAlbumByIdUseCase(albumRepo)
	getAlbumByTitleUseCase := application.NewGetAlbumByTitleUseCase(albumRepo)
	getAlbumByArtistUseCase := application.NewGetAlbumByArtistUseCase(albumRepo)
	getAllAlbumsUseCase := application.NewGetAllAlbumsUseCase(albumRepo)
	updateAlbumUseCase = application.NewUpdateAlbumsUseCase(albumRepo, rb, broadcaster)
	deleteAlbumUseCase := application.NewDeleteAlbumUseCase(albumRepo)

	return &Dependencies{
		AlbumSaveController:        controllers.NewAlbumSaveController(createAlbumUseCase),
		AlbumGetByIdController:     controllers.NewAlbumGetByIdController(getAlbumByIdUseCase),
		AlbumGetByTitleController:  controllers.NewAlbumGetByTitleController(getAlbumByTitleUseCase),
		AlbumGetByArtistController: controllers.NewAlbumGetByArtistController(getAlbumByArtistUseCase),
		AlbumGetAllController:      controllers.NewAlbumGetAllController(getAllAlbumsUseCase),
		AlbumUpdateController:      controllers.NewAlbumUpdateController(updateAlbumUseCase),
		AlbumDeleteController:      controllers.NewAlbumDeleteController(deleteAlbumUseCase),

		ShortPollingStockController: controllers.NewShortPollingStockController(client),
		ShortPollingPriceController: controllers.NewShortPollingPriceController(client),
		LongPollingController:       controllers.NewLongPollingController(client),

		Broadcaster: broadcaster,
	}
}
