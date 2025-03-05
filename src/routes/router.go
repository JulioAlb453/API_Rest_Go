package routes

import (
	albumInfra "API_ejemplo/src/album/infraestructure"
	supplierInfra "API_ejemplo/src/supplier/infraestructure"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine{
	router := gin.Default()
	router.Use(cors.Default()) 
	
	albumDeps := albumInfra.Init()
	albumGroup := router.Group("/albums")
	albumInfra.Routes(albumGroup, albumDeps)

	supplierDeps := supplierInfra.InitSupplierDeps()
	supplierGroup := router.Group("/suppliers")
	supplierInfra.Routes(supplierGroup, supplierDeps)
	
	return router
}