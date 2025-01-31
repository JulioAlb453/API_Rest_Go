package controllers

import (
    "API_ejemplo/src/album/application"
    "API_ejemplo/src/errores"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "github.com/gin-gonic/gin"
    "net/http"
    "errors" 
)

type AlbumDeleteController struct {
    UseCase *application.DeleteAlbumUseCase
}

func NewAlbumDeleteController(useCase *application.DeleteAlbumUseCase) *AlbumDeleteController {
    return &AlbumDeleteController{UseCase: useCase}
}

func (ac *AlbumDeleteController) DeleteAlbumHandler(c *gin.Context) {
    id := c.Param("id")

    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        errores.SendErrorResponse(c, http.StatusBadRequest, errors.New("ID inválido")) 
        return
    }

    _, err = ac.UseCase.Execute(c.Request.Context(), objectId)
    if err != nil {
        errores.SendErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Album eliminado exitosamente", 
    })
}