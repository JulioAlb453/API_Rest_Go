package application

import (
	"API_ejemplo/src/supplier/domain"
	"context"
)

type CreateSupplierUseCase struct {
	repo domain.ISupplier
}

func NewCreateSupplierUseCase(repo domain.ISupplier) *CreateSupplierUseCase {
	return &CreateSupplierUseCase{repo: repo}
}

func (uc *CreateSupplierUseCase) Execute(ctx context.Context, supplier domain.Supplier) error {
	if supplier.Name == "" || supplier.Phone == "" || supplier.Address == "" {
		return domain.ErrMissingFields
	}

	if supplier.Email == "" {
		return domain.ErrInvalidEmail
	}

	if err := uc.repo.Save(ctx, supplier); err != nil {
		return err
	}

	return nil
}
