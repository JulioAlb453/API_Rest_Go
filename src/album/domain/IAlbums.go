package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


type IAlbums interface{
    Save(ctx context.Context, album Album) error
	GetAlbumsById(ctx context.Context, id primitive.ObjectID) (Album, error)
	GetAllAlbums(ctx context.Context) ([]Album, error)
	GetAlbumsByArtist(ctx context.Context, title string) ([]Album, error)
	GetAlbumsByTitle(ctx context.Context, artist string)([]Album, error)
    Update(ctx context.Context, album Album) (Album, error)
	Delete(ctx context.Context, id primitive.ObjectID)  error
}