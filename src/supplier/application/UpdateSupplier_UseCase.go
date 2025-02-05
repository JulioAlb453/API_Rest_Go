package application

import (
	"API_ejemplo/src/supplier/domain"
	"context"
	"errors"
)

type UpdateSupplierUseCase struct {
	repo domain.ISupplier
}

func NewUpdateSupplierUseCase(repo domain.ISupplier) *UpdateSupplierUseCase{
	return &UpdateSupplierUseCase{repo: repo}
}

func (uc *UpdateSupplierUseCase) Execute(ctx context.Context, supplier domain.Supplier) (domain.Supplier, error) {
	if supplier.Name == "" || supplier.Address == "" {
		return domain.Supplier{}, errors.New("Nombre y direccion son requeridos")
	}

	if supplier.Email == "" {
		return domain.Supplier{}, domain.ErrInvalidEmail
	}

	existingSupplier, err := uc.repo.GetSupplierById(ctx, supplier.Id)
	if err != nil {
		if errors.Is(err, domain.ErrSupplerNotFound) {
			return domain.Supplier{}, domain.ErrSupplerNotFound
		}
		return domain.Supplier{}, err
	}

	existingSupplier.Name = supplier.Name
	existingSupplier.Address = supplier.Address
	existingSupplier.Email = supplier.Email
	existingSupplier.Phone = supplier.Phone

	updateSupplier, err := uc.repo.Update(ctx, existingSupplier)
	if err != nil {
		return domain.Supplier{}, err
	}
	return updateSupplier, nil
}
