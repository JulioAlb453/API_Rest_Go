package domain

import (
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Album struct {
	Id        primitive.ObjectID  `bson:"_id,omitempty"`
	Title     string `bson:"Title"`
	Artist    string `bson:"Artist"`
	Year      string `bson:"Year"`
	CreatedAt time.Time `bson:"createdAt"`
}
