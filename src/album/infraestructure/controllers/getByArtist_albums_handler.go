package controllers

import (
    "API_ejemplo/src/album/application"
    "API_ejemplo/src/errores"
    "github.com/gin-gonic/gin"
    "net/http"
)

type AlbumGetByArtistController struct {
    UseCase *application.GetAlbumByArtistUseCase
}

func NewAlbumGetByArtistController(useCase *application.GetAlbumByArtistUseCase) *AlbumGetByArtistController {
    return &AlbumGetByArtistController{UseCase: useCase}
}

func (ac *AlbumGetByArtistController) GetAlbumByArtistHandler(c *gin.Context) {
    artist := c.Param("artist")

    albums, err := ac.UseCase.Execute(c.Request.Context(), artist)
    if err != nil {
        errores.SendErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, albums)
}
