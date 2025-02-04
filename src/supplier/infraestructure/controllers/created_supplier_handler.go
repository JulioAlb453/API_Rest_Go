package controllers

import (
	"API_ejemplo/src/errores"
	"API_ejemplo/src/supplier/application"
	"API_ejemplo/src/supplier/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SupplierSaveController struct {
	UseCase *application.CreateSupplierUseCase
}

func NewSupplierSaveController(useCase *application.CreateSupplierUseCase) *SupplierSaveController {
	return &SupplierSaveController{UseCase: useCase}
}

func (sc *SupplierSaveController) CreateSupplierHandler(c *gin.Context) {
	var supplier domain.Supplier

	if err := c.ShouldBindJSON(&supplier); err != nil {
		errores.SendErrorResponse(c, http.StatusBadRequest, &domain.FieldError{
			Field:   "body",
			Message: "Formato de solicitud inválido " + err.Error(),
		})
		return
	}

	if supplier.Name == "" {
		errores.SendErrorResponse(c, http.StatusBadRequest, &domain.FieldError{
			Field:   "Name",
			Message: "El nombre es obligatorio",
		})
		return
	}

	if supplier.Address == "" {
		errores.SendErrorResponse(c, http.StatusBadRequest, &domain.FieldError{
			Field:   "Address",
			Message: "La dirección es obligatoria",
		})
		return
	}

	if supplier.Email == "" {
		errores.SendErrorResponse(c, http.StatusBadRequest, &domain.FieldError{
			Field:   "Email",
			Message: "El email es obligatorio",
		})
		return
	}

	if supplier.Phone == "" {
		errores.SendErrorResponse(c, http.StatusBadRequest, &domain.FieldError{
			Field:   "Phone",
			Message: "El telefono es obligatorio",
		})
		return
	}

	if err := sc.UseCase.Execute(context.Background(), supplier); err != nil {
		errores.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Proovedor registrado correctamente"})
}
