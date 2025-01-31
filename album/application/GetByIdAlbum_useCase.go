package application

import (
    "API_ejemplo/album/domain"
    "context"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type GetAlbumByIdUseCase struct {
    repo domain.IAlbums
}

func NewGetAlbumByIdUseCase(repo domain.IAlbums) *GetAlbumByIdUseCase {
    return &GetAlbumByIdUseCase{repo: repo}
}

func (uc *GetAlbumByIdUseCase) Execute(ctx context.Context, id primitive.ObjectID) (domain.Album, error) {
    album, err := uc.repo.GetAlbumsById(ctx, id)
    if err != nil {
        return domain.Album{}, err
    }
    return album, nil
}
