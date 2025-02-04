package controllers

import (
	"API_ejemplo/src/errores"
	"API_ejemplo/src/supplier/application"
	"net/http"
	"errors"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SupplierDeleteController struct{
	UseCase *application.DeleteSupplierUseCase
}

func NewSupplierDeleteController(useCase *application.DeleteSupplierUseCase) *SupplierDeleteController{
	return &SupplierDeleteController{UseCase: useCase}
}

func (sc *SupplierDeleteController) DeleteSupplierHandler(c *gin.Context){
	id := c.Param("id")

	objetId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		errores.SendErrorResponse(c, http.StatusBadRequest, errors.New("ID invalido"))
		return
	}

	_, err = sc.UseCase.Execute(c.Request.Context(), objetId)

	if err != nil {
		errores.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	
	c.JSON(http.StatusOK, gin.H{
		"message": "Proovedor eliminado correctamente",
	})

}

