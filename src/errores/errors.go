package errores

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ErrorResponse struct {
	Status string `json:"status"`
	Message string `json:"message"`
}

func SendErrorResponse (c *gin.Context, statusCode int, err error){
	c.JSON(statusCode, ErrorResponse{
		Status: http.StatusText(statusCode),
        Message: err.Error(),
	})
}