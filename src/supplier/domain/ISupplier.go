package domain


import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type ISupplier interface{
	Save(ctx context.Context, supplier Supplier)
	GetSupplierById(ctx context.Context, id primitive.ObjectID)(Supplier, error)
	GetAllSupplier(ctx context.Context) ([]Supplier, error)
	Update (ctx context.Context, supplier Supplier)(Supplier, error)
	Delete (ctx context.Context, id primitive.ObjectID) error
}