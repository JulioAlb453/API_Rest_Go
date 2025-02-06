package infraestructure

import (
	"API_ejemplo/src/album/application"
	"API_ejemplo/src/album/infraestructure/controllers"
	"API_ejemplo/src/album/infraestructure/repository"
	"API_ejemplo/src/core"
	"log"
)

type Dependencies struct {
	AlbumSaveController       *controllers.AlbumSaveController
	AlbumGetByIdController    *controllers.AlbumGetByIdController
	AlbumGetAllController     *controllers.AlbumGetAllController
	AlbumUpdateController     *controllers.AlbumUpdateController
	AlbumDeleteController     *controllers.AlbumDeleteController

	ShortPollingStockController    *controllers.ShortPollingStockController
	LongPollingController     *controllers.LongPollingController
	ShortPollingPriceController    *controllers.ShortPollingPriceController
}

func Init() *Dependencies {
	conn := core.Connect()
	if conn == nil {
		log.Fatal("Error al conectar con la base de datos")
	}
	db := conn.Database("MundyWalk")
	client := conn

	albumRepo := repository.NewMongoAlbumRepository(db)

	createAlbumUseCase := application.NewCreatedAlbumUseCase(albumRepo)
	getAlbumByIdUseCase := application.NewGetAlbumByIdUseCase(albumRepo)
	getAllAlbumsUseCase := application.NewGetAllAlbumsUseCase(albumRepo)
	updateAlbumUseCase := application.NewUpdateAlbumsUseCase(albumRepo)
	deleteAlbumUseCase := application.NewDeleteAlbumUseCase(albumRepo)

	albumSaveController := controllers.NewAlbumSaveController(createAlbumUseCase)
	albumGetByIdController := controllers.NewAlbumGetByIdController(getAlbumByIdUseCase)
	albumGetAllController := controllers.NewAlbumGetAllController(getAllAlbumsUseCase)
	albumUpdateController := controllers.NewAlbumUpdateController(updateAlbumUseCase)
	albumDeleteController := controllers.NewAlbumDeleteController(deleteAlbumUseCase)

	shortPollingController := controllers.NewShortPollingStockController(client)
	shortPollingPriceController := controllers.NewShortPollingPriceController(client)
	longPollingController := controllers.NewLongPollingController(client)

	return &Dependencies{
		AlbumSaveController:       albumSaveController,
		AlbumGetByIdController:    albumGetByIdController,
		AlbumGetAllController:     albumGetAllController,
		AlbumUpdateController:     albumUpdateController,
		AlbumDeleteController:     albumDeleteController,
		ShortPollingStockController: 	shortPollingController,
		ShortPollingPriceController:  shortPollingPriceController,
		LongPollingController: 		longPollingController,
	}
}
