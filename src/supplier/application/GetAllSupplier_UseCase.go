package application

import (
	"API_ejemplo/src/supplier/domain"
	"context"
)

type GetAllSupplierUseCase struct {
	repo domain.ISupplier
}

func NewGetAllSupplierUseCase(repo domain.ISupplier) *GetAllSupplierUseCase {
	return &GetAllSupplierUseCase{repo: repo}
}

func (uc *GetAllSupplierUseCase) Execute(ctx context.Context) ([]domain.Supplier, error) {
	return uc.repo.GetAllSupplier(ctx)
}
