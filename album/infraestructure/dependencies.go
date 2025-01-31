package infraestructure

import (
	"API_ejemplo/album/application"
	"API_ejemplo/album/infraestructure/controllers"
	"API_ejemplo/album/infraestructure/repository"
	"API_ejemplo/core"
	"log"
)

type Dependencies struct {
	AlbumSaveController       *controllers.AlbumSaveController
	AlbumGetByIdController    *controllers.AlbumGetByIdController
	AlbumGetAllController     *controllers.AlbumGetAllController
	AlbumUpdateController     *controllers.AlbumUpdateController
	AlbumDeleteController     *controllers.AlbumDeleteController
}

func Init() *Dependencies {
	conn := core.Connect()
	if conn == nil {
		log.Fatal("Error al conectar con la base de datos")
	}
	db := conn.Database("MundyWalk")

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

	return &Dependencies{
		AlbumSaveController:       albumSaveController,
		AlbumGetByIdController:    albumGetByIdController,
		AlbumGetAllController:     albumGetAllController,
		AlbumUpdateController:     albumUpdateController,
		AlbumDeleteController:     albumDeleteController,
	}
}
