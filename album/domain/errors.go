package domain

import "errors"

var (
	ErrMissingFields = errors.New("faltan campos obligatorios")
	ErrInvalidData   = errors.New("datos inválidos")
	ErrMissingFieldsWithError = errors.New("datos inválidos")
	ErrInternalError = errors.New("ocurrió un error interno al procesar la solicitud")
	ErrInvalidInput  = errors.New("los datos enviados no son válidos")
	ErrInvalidYear = errors.New("Año no valido")
	ErrAlbumNotFound = errors.New("Album no encontrado")
)

