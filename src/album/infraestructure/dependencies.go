package infraestructure

import (
	"API_ejemplo/src/album/application"
	"API_ejemplo/src/album/infraestructure/controllers"
	"API_ejemplo/src/album/infraestructure/repository"
	"API_ejemplo/src/core"
	"API_ejemplo/src/shared/broker"
	"log"
)

type Dependencies struct {
	AlbumSaveController          *controllers.AlbumSaveController
	AlbumGetByIdController       *controllers.AlbumGetByIdController	
	AlbumGetByArtistController   *controllers.AlbumGetByArtistController 
	AlbumGetByTitleController    *controllers.AlbumGetByTitleController
	AlbumGetAllController        *controllers.AlbumGetAllController
	AlbumUpdateController        *controllers.AlbumUpdateController
	AlbumDeleteController        *controllers.AlbumDeleteController

	ShortPollingStockController  *controllers.ShortPollingStockController
	LongPollingController        *controllers.LongPollingController
	ShortPollingPriceController  *controllers.ShortPollingPriceController
}

func Init() *Dependencies {
	conn := core.Connect()
	if conn == nil {
		log.Fatal("Error al conectar con la base de datos")
	}
	db := conn.Database("MundyWalk")
	client := conn

	rb, err := broker.NewRabbitMQBroker("amqp://guest:guest@3.209.113.62:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	albumRepo := repository.NewMongoAlbumRepository(db)

	consumeAlbumEventUseCase := application.NewConsumeAlbumEventUseCase(albumRepo)

	err = rb.Consume("album_queue", "album.data", func(msg []byte) error {
		log.Println("🎧 Mensaje recibido desde RabbitMQ:", string(msg))
		return consumeAlbumEventUseCase.Handle(msg)
	})
	if err != nil {
		log.Fatalf("Error al consumir de RabbitMQ: %v", err)
	}
	

	createAlbumUseCase := application.NewCreatedAlbumUseCase(albumRepo)
	getAlbumByIdUseCase := application.NewGetAlbumByIdUseCase(albumRepo)
	getAlbumByTitleUseCase := application.NewGetAlbumByTitleUseCase(albumRepo) 
	getAlbumByArtistUseCase := application.NewGetAlbumByArtistUseCase(albumRepo) 
	getAllAlbumsUseCase := application.NewGetAllAlbumsUseCase(albumRepo)
	updateAlbumUseCase := application.NewUpdateAlbumsUseCase(albumRepo, rb)
	deleteAlbumUseCase := application.NewDeleteAlbumUseCase(albumRepo)

	albumSaveController := controllers.NewAlbumSaveController(createAlbumUseCase)
	albumGetByIdController := controllers.NewAlbumGetByIdController(getAlbumByIdUseCase)
	albumGetByTitleController := controllers.NewAlbumGetByTitleController(getAlbumByTitleUseCase) 
	albumGetByArtistController := controllers.NewAlbumGetByArtistController(getAlbumByArtistUseCase) 
	albumGetAllController := controllers.NewAlbumGetAllController(getAllAlbumsUseCase)
	albumUpdateController := controllers.NewAlbumUpdateController(updateAlbumUseCase)
	albumDeleteController := controllers.NewAlbumDeleteController(deleteAlbumUseCase)

	shortPollingController := controllers.NewShortPollingStockController(client)
	shortPollingPriceController := controllers.NewShortPollingPriceController(client)
	longPollingController := controllers.NewLongPollingController(client)

	return &Dependencies{
		AlbumSaveController:          albumSaveController,
		AlbumGetByIdController:       albumGetByIdController,
		AlbumGetByTitleController:    albumGetByTitleController,  
		AlbumGetByArtistController:   albumGetByArtistController, 
		AlbumGetAllController:        albumGetAllController,
		AlbumUpdateController:        albumUpdateController,
		AlbumDeleteController:        albumDeleteController,
		ShortPollingStockController:  shortPollingController,
		ShortPollingPriceController:  shortPollingPriceController,
		LongPollingController:        longPollingController,
	}
}
