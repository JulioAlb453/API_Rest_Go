package application

import (
    "API_ejemplo/src/album/domain"
    "context"
    "errors"
)

type  GetAlbumByArtistUseCase struct {
    repo domain.IAlbums
}

func NewGetAlbumByArtistUseCase(repo domain.IAlbums) *GetAlbumByArtistUseCase {
    return &GetAlbumByArtistUseCase{repo: repo}
}

func (uc *GetAlbumByArtistUseCase) Execute(ctx context.Context, artist string) ([]domain.Album, error) {
    albums, err := uc.repo.GetAlbumsByArtist(ctx, artist)
    if err != nil {
        return nil, errors.New("Error al obtener álbumes por artista: " + err.Error())
    }
    return albums, nil
}
