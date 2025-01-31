package application

import (
	"API_ejemplo/src/album/domain"
	"context"
)

type CreateAlbumUseCase struct {
	repo domain.IAlbums
}

func NewCreatedAlbumUseCase(repo domain.IAlbums) *CreateAlbumUseCase {
	return &CreateAlbumUseCase{repo: repo}
}

func (uc *CreateAlbumUseCase) Execute(ctx context.Context, album domain.Album) error {
	// Validar campos obligatorios
	if album.Artist == "" || album.Title == ""  {
		return domain.ErrMissingFields
	}

	if album.Year == "" {
		return domain.ErrInvalidYear
	}

	if err := uc.repo.Save(ctx, album); err != nil {
		return err 
	}

	return nil
}
