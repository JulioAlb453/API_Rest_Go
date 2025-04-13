package infraestructure

import (
	"API_ejemplo/src/core"
	"API_ejemplo/src/supplier/application"
	"API_ejemplo/src/supplier/infraestructure/controllers"
	"API_ejemplo/src/supplier/infraestructure/notification"
	"API_ejemplo/src/supplier/infraestructure/repository"
	"log"
)

// Dependencies contiene todos los controladores inyectados
type Dependencies struct {
	SupplierSaveController     *controllers.SupplierSaveController
	SupplierGetByIdController  *controllers.SupplierGetByIdController
	SupplierGetAllController   *controllers.SupplierGetAllController
	SupplierUpdateController   *controllers.SupplierUpdateController
	SupplierDeleteController   *controllers.SupplierDeleteController
}

// InitSupplierDeps inicializa y retorna todas las dependencias inyectadas
func InitSupplierDeps() *Dependencies {
	// 1. Establecer conexión con MongoDB
	client := core.Connect()
	if client == nil {
		log.Fatal("❌ Error fatal: No se pudo conectar a la base de datos")
	}
	db := client.Database("MundyWalk")

	supplierRepo := repository.NewMongoSupplierRepository(db)

	emailNotifier := &notification.FakeEmailSender{}

	createUC := application.NewCreateSupplierUseCase(supplierRepo, emailNotifier)
	getByIdUC := application.NewGetSupplierByIdUSeCase(supplierRepo)
	getAllUC := application.NewGetAllSupplierUseCase(supplierRepo)
	updateUC := application.NewUpdateSupplierUseCase(supplierRepo)
	deleteUC := application.NewDeleteSupplierUseCase(supplierRepo)

	// 5. Inicializar los controladores con los casos de uso
	return &Dependencies{
		SupplierSaveController:    controllers.NewSupplierSaveController(createUC),
		SupplierGetByIdController: controllers.NewSupplierGetByIdController(getByIdUC),
		SupplierGetAllController:  controllers.NewSpplierGetAllController(getAllUC),
		SupplierUpdateController:  controllers.NewSupplierUpdateController(updateUC),
		SupplierDeleteController: controllers.NewSupplierDeleteController(deleteUC),
	}
}