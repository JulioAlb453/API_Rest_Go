package infraestructure

import "github.com/gin-gonic/gin"

func Routes(group *gin.RouterGroup, deps *Dependencies) {
	group.POST("/", deps.SupplierSaveController.CreateSupplierHandler)
	group.GET("/", deps.SupplierGetAllController.GetAllSupplierHandler)
	group.GET("/:id", deps.SupplierGetByIdController.GetSupplierByIdHandler)
	group.PUT("/:id", deps.SupplierUpdateController.UpdateSupplierHandler)
	group.DELETE("/:id", deps.SupplierDeleteController.DeleteSupplierHandler)
}
