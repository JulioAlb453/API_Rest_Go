package controllers

import (
    "API_ejemplo/album/application"
    "API_ejemplo/errores"
    "API_ejemplo/album/domain"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "net/http"
    "errors"  
)

type AlbumUpdateController struct {
    UseCase *application.UpdateAlbumsUseCase
}

func NewAlbumUpdateController(useCase *application.UpdateAlbumsUseCase) *AlbumUpdateController {
    return &AlbumUpdateController{UseCase: useCase}
}

func (ac *AlbumUpdateController) UpdateAlbumHandler(c *gin.Context) {
    id := c.Param("id")
    var album domain.Album

    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        errores.SendErrorResponse(c, http.StatusBadRequest, errors.New("ID inválido")) // Crear error con errors.New
        return
    }

    if err := c.ShouldBindJSON(&album); err != nil {
        errores.SendErrorResponse(c, http.StatusBadRequest, domain.ErrInvalidData)
        return
    }

    album.Id = objectId 

    updatedAlbum, err := ac.UseCase.Execute(c.Request.Context(), album)
    if err != nil {
        errores.SendErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Album actualizado exitosamente",
        "album": updatedAlbum,
    })
}
