package controllers

import (
    "API_ejemplo/src/album/application"
    "API_ejemplo/src/errores"
    "github.com/gin-gonic/gin"
    "net/http"
)

type AlbumGetByTitleController struct {
    UseCase *application.GetAlbumByTitleUseCase
}

func NewAlbumGetByTitleController(useCase *application.GetAlbumByTitleUseCase) *AlbumGetByTitleController {
    return &AlbumGetByTitleController{UseCase: useCase}
}

func (ac *AlbumGetByTitleController) GetAlbumByTitleHandler(c *gin.Context) {
    title := c.Param("title")

    albums, err := ac.UseCase.Execute(c.Request.Context(), title)
    if err != nil {
        errores.SendErrorResponse(c, http.StatusInternalServerError, err)
        return
    }

    c.JSON(http.StatusOK, albums)
}
