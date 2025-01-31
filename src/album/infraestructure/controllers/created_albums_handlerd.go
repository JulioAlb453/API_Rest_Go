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
        errores.SendErrorResponse(c, http.StatusBadRequest, &domain.FieldError{
            Field:   "body",
            Message: "Formato de solicitud inválido: " + err.Error(),
        })
        return
    }

    if album.Title == "" {
        errores.SendErrorResponse(c, http.StatusBadRequest, &domain.FieldError{
            Field:   "Title",
            Message: "El título es obligatorio",
        })
        return
    }

    if album.Artist == "" {
        errores.SendErrorResponse(c, http.StatusBadRequest, &domain.FieldError{
            Field:   "Artist",
            Message: "El artista es obligatorio",
        })
        return
    }

    if album.Price <= 0 {
        errores.SendErrorResponse(c, http.StatusBadRequest, &domain.FieldError{
            Field:   "Price",
            Message: "El precio debe ser mayor que 0",
        })
        return
    }

    if album.Stock <= 0 {
        errores.SendErrorResponse(c, http.StatusBadRequest, &domain.FieldError{
            Field:   "Stock",
            Message: "El stock debe ser mayor que 0",
        })
        return
    }

    if err := ac.UseCase.Execute(context.Background(), album); err != nil {
        errores.SendErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    // Responder con el mensaje de éxito si todo fue correcto
    c.JSON(http.StatusOK, gin.H{"message": "Álbum creado correctamente"})
}