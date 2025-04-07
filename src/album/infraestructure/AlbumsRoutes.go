package infraestructure

import "github.com/gin-gonic/gin"

func Routes(group *gin.RouterGroup, deps *Dependencies) {
	group.POST("/", deps.AlbumSaveController.CreateAlbumHandler)
	group.GET("/", deps.AlbumGetAllController.GetAllAlbumsHandler)
	group.GET("/:id", deps.AlbumGetByIdController.GetAlbumByIdHandler)
	group.GET("/search/title/:title", deps.AlbumGetByTitleController.GetAlbumByTitleHandler)  
	group.GET("/search/artist/:artist", deps.AlbumGetByArtistController.GetAlbumByArtistHandler)
	group.PUT("/:id", deps.AlbumUpdateController.UpdateAlbumHandler)
	group.DELETE("/:id", deps.AlbumDeleteController.DeleteAlbumHandler)

	group.GET("/short-polling-stock", deps.ShortPollingStockController.ShortPollingStockHandler)
	group.GET("/short-polling-price", deps.ShortPollingPriceController.ShortPollingPriceHandler)
	group.GET("/long-polling", deps.LongPollingController.LongPollingHandler)

}
