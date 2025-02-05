package infraestructure

import (
	"API_ejemplo/src/core"
	"API_ejemplo/src/supplier/application"
	"API_ejemplo/src/supplier/infraestructure/controllers"
	"API_ejemplo/src/supplier/infraestructure/repository"
	"log"
)

type Dependencies struct {
	SupplierSaveController  		*controllers.SupplierSaveController
	SupplierGetByIdController  		*controllers.SupplierGetByIdController
	SupplierGetAllController 		*controllers.SupplierGetAllController
	SupplierUpdateController 		*controllers.SupplierUpdateController
	SupplierDeleteController 		*controllers.SupplierDeleteController
}

func Init() *Dependencies{
	conn := core.Connect()

	if conn == nil {
		log.Fatal("Error al conectar la base de datos")
	}

	db := conn.Database("MundyWalk")

	supplierRepo := repository.NewMongoSupplierRepository(db)

	createSupplierUseCase := application.NewCreateSupplierUseCase(supplierRepo)
	getSupplierByIdUseCase := application.NewGetSupplierByIdUSeCase(supplierRepo)
	getAllSupplierUseCase := application.NewGetAllSupplierUseCase(supplierRepo)
	updateSupplierUseCase := application.NewUpdateSupplierUseCase(supplierRepo)
	deleteSupplierUseCase := application.NewDeleteSupplierUseCase(supplierRepo)	


	supplierSaveController := controllers.NewSupplierSaveController(createSupplierUseCase)
	supplierGetByIdController := controllers.NewSupplierGetByIdController(getSupplierByIdUseCase)
	supplierGetAllController := controllers.NewSpplierGetAllController(getAllSupplierUseCase)
	supplierUpdateController := controllers.NewSupplierUpdateController(updateSupplierUseCase)
	supplierDeleteController := controllers.NewSupplierDeleteController(deleteSupplierUseCase)

	return &Dependencies{
		SupplierSaveController: supplierSaveController,
		SupplierGetAllController: supplierGetAllController,
		SupplierGetByIdController: supplierGetByIdController,
		SupplierUpdateController: supplierUpdateController,
		SupplierDeleteController: supplierDeleteController,
	}

}

