package application

import (
	"API_ejemplo/src/album/domain"
	"context"
)

type GetAllAlbumsUseCase struct{
	repo domain.IAlbums
}

func NewGetAllAlbumsUseCase(repo domain.IAlbums) *GetAllAlbumsUseCase{
    return &GetAllAlbumsUseCase{repo: repo}
}

func (uc *GetAllAlbumsUseCase) Execute(ctx context.Context)([]domain.Album, error){
	return uc.repo.GetAllAlbums(ctx)
}
