package domain

import (
	"errors"
	"fmt"
)

type FieldError struct {
	Field   string
	Message string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("campo: '%s', error: '%s'", e.Field, e.Message)
}

var (
	ErrMissingFields          = errors.New("Faltan campos obligatorios")
	ErrInvalidData            = errors.New("Datos inválidos")
	ErrMissingFieldsWithError = errors.New("Datos inválidos")
	ErrInternalError          = errors.New("Ocurrió un error interno al procesar la solicitud")
	ErrInvalidInput           = errors.New("Los datos enviados no son válidos")
	ErrInvalidEmail           = errors.New("Email no valido")
	ErrSupplerNotFound        = errors.New("Prooverdor no encontrado")
)
