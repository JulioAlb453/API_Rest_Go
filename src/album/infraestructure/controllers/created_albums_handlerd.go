package controllers

import (
	"API_ejemplo/src/album/application"
	"API_ejemplo/src/album/domain"
	"API_ejemplo/src/errores"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumSaveController struct {
	UseCase *application.CreateAlbumUseCase
}

func NewAlbumSaveController(useCase *application.CreateAlbumUseCase) *AlbumSaveController {
	return &AlbumSaveController{UseCase: useCase}
}

func (ac *AlbumSaveController) CreateAlbumHandler(c *gin.Context) {
	var album domain.Album

	if err := c.ShouldBindJSON(&album); err != nil {
		errores.SendErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	if album.Title == "" || album.Artist == "" {
		errores.SendErrorResponse(c, http.StatusBadRequest, domain.ErrMissingFields)
		return
	}

	if err := ac.UseCase.Execute(context.Background(), album); err != nil {
		errores.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Álbum creado correctamente"})
}
