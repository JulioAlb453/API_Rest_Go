package domain


import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Supplier struct{

	Id primitive.ObjectID `bson:"_id,omitempty"`
	Name string `bson:"Name"`
	Phone string `bson:"Phone"`
	Email string `bson:"Email"`
	Address string `bson:"Address"`	
}