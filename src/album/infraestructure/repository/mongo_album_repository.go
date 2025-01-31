package repository

import (
	"API_ejemplo/src/album/domain"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoAlbumRepository struct {
	db *mongo.Database
}

func NewMongoAlbumRepository(conn *mongo.Database) *MongoAlbumRepository {
	return &MongoAlbumRepository{db: conn}
}

func (r *MongoAlbumRepository) Save(ctx context.Context, album domain.Album) error {
	collection := r.db.Collection("albums")
	album.CreatedAt = time.Now()
	_, err := collection.InsertOne(ctx, bson.M{
		"Title":     album.Title,
		"Artist":    album.Artist,
		"Year":      album.Year,
		"createdAt": album.CreatedAt,
	})
	if err != nil {
		return errors.New("Error al guardar el álbum: " + err.Error())
	}
	return nil
}

func (r *MongoAlbumRepository) GetAlbumsById(ctx context.Context, id primitive.ObjectID) (domain.Album, error) {
	collection := r.db.Collection("albums")

filter := bson.M{"_id": id}

	var album domain.Album

	err := collection.FindOne(ctx, filter).Decode(&album)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return domain.Album{}, domain.ErrAlbumNotFound
		}
		return domain.Album{}, errors.New("error al buscar el álbum: " + err.Error())
	}

	return album, nil
}

func (r *MongoAlbumRepository) GetAllAlbums(ctx context.Context) ([]domain.Album, error) {
	collection := r.db.Collection("albums")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, errors.New("Error al obtener los álbumes: " + err.Error())
	}
	defer cursor.Close(ctx)

	var albums []domain.Album
	for cursor.Next(ctx) {
		var album domain.Album
		if err := cursor.Decode(&album); err != nil {
			return nil, errors.New("Error al decodificar un álbum: " + err.Error())
		}
		albums = append(albums, album)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.New("Error iterando sobre los álbumes: " + err.Error())
	}

	return albums, nil
}

func (r *MongoAlbumRepository) Update(ctx context.Context, album domain.Album) (domain.Album, error) {
	collection := r.db.Collection("albums")
	objectId, err := primitive.ObjectIDFromHex(album.Id.Hex())
	 if err !=nil {
		return domain.Album{}, errors.New("ID invalido")
	 }
	
	filter := bson.M{"_id": objectId}
	update := bson.M{
		"$set": bson.M{
			"Title":    album.Title,
			"Artist":   album.Artist,
			"Year":     album.Year,
			"updatedAt": time.Now(),
		},
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil{
		return domain.Album{}, errors.New("Error al actualizar el album" + err.Error())
	}
	if result.MatchedCount == 0{
		return domain.Album{}, errors.New("Ningun album coincide el filtro")
	}
	return album, nil
}



func (r *MongoAlbumRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
    collection := r.db.Collection("albums")

    result, err := collection.DeleteOne(ctx, bson.M{"_id": id})
    if err != nil {
        return err 
    }
    if result.DeletedCount == 0 {
        return domain.ErrAlbumNotFound 
    }
    return nil 
}
