package infraestructure

import "github.com/gin-gonic/gin"

func Routes(deps *Dependencies) *gin.Engine{
	router := gin.Default()

	router.POST("/supplier", deps.SupplierSaveController.CreateSupplierHandler)
	router.GET("/supplier", deps.SupplierGetAllController.GetAllSupplierHandler)
	router.GET("/supplier/", deps.SupplierGetByIdController.GetSupplierByIdHandler)
	router.PUT("/supplier/", deps.SupplierUpdateController.UpdateSupplierHandler)
	router.DELETE("/supplier/", deps.SupplierDeleteController.DeleteSupplierHandler)
	
	return router
}