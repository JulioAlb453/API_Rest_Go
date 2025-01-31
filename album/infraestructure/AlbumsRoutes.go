package infraestructure

import (
	"github.com/gin-gonic/gin"
)

func Routes (deps *Dependencies) *gin.Engine{
	router := gin.Default()

   
	router.POST("/albums", deps.AlbumSaveController.CreateAlbumHandler)
	router.GET("/albums/:id", deps.AlbumGetByIdController.GetAlbumHandler)
	router.GET("/albums", deps.AlbumGetAllController.GetAllAlbumsHandler)
	router.PUT("/albums/:id", deps.AlbumUpdateController.UpdateAlbumHandler)
	router.DELETE("/albums/:id", deps.AlbumDeleteController.DeleteAlbumHandler)

	return router
}