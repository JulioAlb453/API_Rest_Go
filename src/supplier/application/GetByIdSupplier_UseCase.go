package application

import (
	"API_ejemplo/src/supplier/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type GetSupplierByIdSUseCase struct{
	repo domain.ISupplier
}

func NewGetSupplierByIdUSeCase(repo domain.ISupplier) *GetSupplierByIdSUseCase{
	return &GetSupplierByIdSUseCase{repo: repo}
}

func (uc *GetSupplierByIdSUseCase) Execute(ctx context.Context, id primitive.ObjectID) (domain.Supplier, error){
	supplier, err := uc.repo.GetSupplierById(ctx, id)
	if err != nil{
		return domain.Supplier{}, err
	}
	return supplier, nil
}