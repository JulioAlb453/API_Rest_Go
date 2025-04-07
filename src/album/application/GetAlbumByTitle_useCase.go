package application

import (
	"API_ejemplo/src/album/domain"
	"context"
	"errors"
)

type GetAlbumByTitleUseCase struct {
	repo domain.IAlbums
}

func NewGetAlbumByTitleUseCase(repo domain.IAlbums) *GetAlbumByTitleUseCase{
	return &GetAlbumByTitleUseCase{repo: repo}
}

func (uc *GetAlbumByTitleUseCase) Execute(ctx context.Context, title string)([] domain.Album, error){
	albums, err := uc.repo.GetAlbumsByTitle(ctx, title)
	if err != nil{
		return nil, errors.New("Error al obtener álbumes por artista: " + err.Error())

	}
	return albums, nil
}