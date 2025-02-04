package controllers

import (
	"API_ejemplo/src/errores"
	"API_ejemplo/src/supplier/application"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SupplierGetByIdController struct {
	UseCase *application.GetSupplierByIdSUseCase
}

func NewSupplierGetByIdController(useCase *application.GetSupplierByIdSUseCase) *SupplierGetByIdController {
	return &SupplierGetByIdController{UseCase: useCase}
}

func (sc *SupplierGetByIdController) GetSupplierByIdHandler(c *gin.Context) {
	id := c.Param("id")
	ObjectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		errores.SendErrorResponse(c, http.StatusBadRequest, errors.New("ID invalido"))
		return
	}

	supplier, err := sc.UseCase.Execute(c.Request.Context(), ObjectId)

	if err != nil {
		errores.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, supplier)
}
