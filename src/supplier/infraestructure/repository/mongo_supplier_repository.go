package repository

import (
	"API_ejemplo/src/supplier/domain"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoSupplierRepository struct {
	db *mongo.Database
}

func NewMongoSupplierRepository(conn *mongo.Database) *MongoSupplierRepository {
	return &MongoSupplierRepository{db: conn}
}

func (r *MongoSupplierRepository) Save(ctx context.Context, supplier domain.Supplier) error {
	collection := r.db.Collection("suppliers")
	_, err := collection.InsertOne(ctx, bson.M{
		"Name":    supplier.Name,
		"Phone":   supplier.Phone,
		"Email":   supplier.Email,
		"Address": supplier.Address,
	})
	if err != nil {
		return errors.New("Error al guardar el proovedor" + err.Error())
	}
	return nil
}

func (r *MongoSupplierRepository) GetSupplierById(ctx context.Context, id primitive.ObjectID) (domain.Supplier, error) {
	collection := r.db.Collection("suppliers")

	filter := bson.M{"_id": id}

	var supplier domain.Supplier

	err := collection.FindOne(ctx, filter).Decode(&supplier)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Supplier{}, domain.ErrSupplerNotFound
		}
		return domain.Supplier{}, errors.New("Erro al buscar el álbum: " + err.Error())
	}

	return supplier, nil
}

func (r *MongoSupplierRepository) GetAllSupplier(ctx context.Context) ([]domain.Supplier, error) {
	collection := r.db.Collection("supplier")
	cursor, err := collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, errors.New("Error al obtener los álbumes: " + err.Error())
	}
	defer cursor.Close(ctx)

	var suppliers []domain.Supplier

	for cursor.Next(ctx) {
		var supplier domain.Supplier
		if err := cursor.Decode(&supplier); err != nil {
			return nil, errors.New("Error al codificar el proovedor: " + err.Error())
		}
		suppliers = append(suppliers, supplier)
	}
	if err := cursor.Err(); err != nil {
		return nil, errors.New("Error iterando a los proovedores: " + err.Error())
	}
	return suppliers, nil
}	

func (r *MongoSupplierRepository) Update(ctx context.Context, supplier domain.Supplier) (domain.Supplier, error){
	collection := r.db.Collection("suppliers")

	objetId, err := primitive.ObjectIDFromHex(supplier.Id.Hex())

	if err != nil {
		return domain.Supplier{}, errors.New("ID invalido")
	}

	filter := bson.M{"_id": objetId}
	update := bson.M{
		"$set": bson.M{

			"Name": supplier.Name,
			"Phone": supplier.Phone,
			"Email": supplier.Email,
			"Address": supplier.Address,
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil{
		return domain.Supplier{}, errors.New("Error al actualizar al proovedor: " + err.Error())
	}

	if result.MatchedCount == 0 {
		return domain.Supplier{}, errors.New("Ningun proovedor coincide con el filtro")
	}
	return supplier, nil
}

func (r *MongoSupplierRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	collection := r.db.Collection("supplier")

	result, err := collection.DeleteOne(ctx, bson.M{"_id": id})

	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return domain.ErrSupplerNotFound
	}

	return nil
}