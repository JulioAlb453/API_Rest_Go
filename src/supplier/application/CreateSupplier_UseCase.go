	package application

	import (
		"API_ejemplo/src/supplier/domain"
		"context"
		"errors"
	)

	type CreateSupplierUseCase struct {
		repo     domain.ISupplier
		notifier domain.EmailNotifier
		}       

	func NewCreateSupplierUseCase(
		repo domain.ISupplier, 
		notifier domain.EmailNotifier,
	) *CreateSupplierUseCase {
		return &CreateSupplierUseCase{
			repo:     repo,
			notifier: notifier,
		}
	}

	func (uc *CreateSupplierUseCase) Execute(ctx context.Context, supplier domain.Supplier) error {
		if supplier.Name == "" || supplier.Phone == "" || supplier.Address == "" {
			return errors.New("campos requeridos faltantes: nombre, teléfono y dirección")
		}

		if supplier.Email == "" {
			return errors.New("email no válido")
		}

		if err := uc.repo.Save(ctx, supplier); err != nil {
			return errors.New("error al guardar el proveedor: " + err.Error())
		}

		subject := "Registro de proveedor exitoso"
		body := "Hola " + supplier.Name + ", gracias por registrarte como proveedor."
		
		if err := uc.notifier.SendEmail(supplier.Email, subject, body); err != nil {
			return errors.New("error al enviar notificación: " + err.Error())
		}

		return nil
	}