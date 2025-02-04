package application

import (
	"API_ejemplo/src/supplier/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteSupplierUseCase struct{
	repo domain.ISupplier
}

func NewDeleteSupplierUseCase(repo domain.ISupplier) *DeleteSupplierUseCase{
	return &DeleteSupplierUseCase{repo: repo}
}

func (uc *DeleteSupplierUseCase) Execute (ctx context.Context, id primitive.ObjectID) (domain.Supplier, error){
	err := uc.repo.Delete(ctx, id)
	if err != nil {
		return domain.Supplier{}, err
	}
	return domain.Supplier{}, nil
}

