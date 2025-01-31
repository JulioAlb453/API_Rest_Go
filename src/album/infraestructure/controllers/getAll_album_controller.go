package controllers

import (
	"API_ejemplo/src/errores"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"API_ejemplo/src/album/application")

type AlbumGetAllController struct {
	UseCase *application.GetAllAlbumsUseCase
}

func NewAlbumGetAllController(useCase *application.GetAllAlbumsUseCase) *AlbumGetAllController{
	return &AlbumGetAllController{UseCase : useCase}
}

func (ac *AlbumGetAllController) GetAllAlbumsHandler(c *gin.Context){
	albums, err := ac.UseCase.Execute(context.Background())

	if err != nil{
		errores.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, albums)
}