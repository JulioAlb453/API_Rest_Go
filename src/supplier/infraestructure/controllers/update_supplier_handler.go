package controllers

import (
	"API_ejemplo/src/errores"
	"API_ejemplo/src/supplier/application"
	"API_ejemplo/src/supplier/domain"
	"errors"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SupplierUpdateController struct {
	UseCase *application.UpdateSupplierUseCase
}

func NewSupplierUpdateController(useCase *application.UpdateSupplierUseCase) *SupplierUpdateController {
	return &SupplierUpdateController{UseCase: useCase}
}

func (sc *SupplierUpdateController) UpdateSupplierHandler(c *gin.Context) {
	id := c.Param("id")

	var supplier domain.Supplier
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		errores.SendErrorResponse(c, http.StatusBadRequest, errors.New("ID inválido"))
	}

	if err := c.ShouldBindJSON(&supplier); err != nil {
		errores.SendErrorResponse(c, http.StatusBadRequest, domain.ErrInvalidData)
		return
	}

	supplier.Id = objectId

	updateSupplier, err := sc.UseCase.Execute(c.Request.Context(), supplier)

	if err != nil {
		errores.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Proovedor actualizado exitosamente",
		"supplier": updateSupplier,
	})
}
