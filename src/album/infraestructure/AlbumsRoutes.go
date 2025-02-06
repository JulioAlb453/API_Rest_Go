package infraestructure

import "github.com/gin-gonic/gin"

func Routes(group *gin.RouterGroup, deps *Dependencies) {
	group.POST("/", deps.AlbumSaveController.CreateAlbumHandler)
	group.GET("/", deps.AlbumGetAllController.GetAllAlbumsHandler)
	group.GET("/:id", deps.AlbumGetByIdController.GetAlbumByIdHandler)
	group.PUT("/:id", deps.AlbumUpdateController.UpdateAlbumHandler)
	group.DELETE("/:id", deps.AlbumDeleteController.DeleteAlbumHandler)
}
