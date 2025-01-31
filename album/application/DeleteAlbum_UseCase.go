package application

import (
	"API_ejemplo/album/domain"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeleteAlbumUseCase struct {
	repo domain.IAlbums
}

func NewDeleteAlbumUseCase(repo domain.IAlbums) *DeleteAlbumUseCase {
	return &DeleteAlbumUseCase{repo: repo}
}

func (uc *DeleteAlbumUseCase) Execute(ctx context.Context, id primitive.ObjectID) (domain.Album, error) {
    err := uc.repo.Delete(ctx, id)
    if err != nil {
        return domain.Album{}, err  
    }
    return domain.Album{}, nil 
}