package controllers

import (
    "API_ejemplo/src/album/application"
    "API_ejemplo/src/errores"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson/primitive"
    "net/http"
    "errors" )

type AlbumGetByIdController struct {
    UseCase *application.GetAlbumByIdUseCase
}

func NewAlbumGetByIdController(useCase *application.GetAlbumByIdUseCase) *AlbumGetByIdController {
    return &AlbumGetByIdController{UseCase: useCase}
}

func (ac *AlbumGetByIdController) GetAlbumByIdHandler(c *gin.Context) {
    id := c.Param("id")
    objectId, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        errores.SendErrorResponse(c, http.StatusBadRequest, errors.New("ID inválido"))
        return
    }

    album, err := ac.UseCase.Execute(c.Request.Context(), objectId)
    if err != nil {
        errores.SendErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, album)
}
