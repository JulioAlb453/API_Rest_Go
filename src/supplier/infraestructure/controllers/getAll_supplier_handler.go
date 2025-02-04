package controllers

import (
	"API_ejemplo/src/errores"
	"API_ejemplo/src/supplier/application"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SupplierGetAllController struct {
	UseCase *application.GetAllSupplierUseCase
}

func NewSpplierGetAllController(useCase *application.GetAllSupplierUseCase) *SupplierGetAllController {
	return &SupplierGetAllController{UseCase: useCase}
}

func (sp *SupplierGetAllController) GetAllSupplierHandler(c *gin.Context) {
	supplier, err := sp.UseCase.Execute(context.Background())

	if err != nil {
		errores.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, supplier)
}
