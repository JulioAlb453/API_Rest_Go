package application

import (
	"API_ejemplo/album/domain"
	"context"
	"errors"
)

type UpdateAlbumsUseCase struct {
	repo domain.IAlbums
}

func NewUpdateAlbumsUseCase(repo domain.IAlbums) *UpdateAlbumsUseCase {
	return &UpdateAlbumsUseCase{repo: repo}
}

func (uc *UpdateAlbumsUseCase) Execute(ctx context.Context, album domain.Album) (domain.Album, error) {
	if album.Artist == "" || album.Title == "" {
		return domain.Album{}, errors.New("invalid input: Artista y Titulo son requeridos")
	}

	existingAlbum, err := uc.repo.GetAlbumsById(ctx, album.Id)
	if err != nil {
		if errors.Is(err, domain.ErrAlbumNotFound) {
			return domain.Album{}, domain.ErrAlbumNotFound
		}
		return domain.Album{}, err
	}

	existingAlbum.Title = album.Title
	existingAlbum.Artist = album.Artist
	existingAlbum.Year = album.Year

	updatedAlbum, err := uc.repo.Update(ctx, existingAlbum)
	if err != nil {
		return domain.Album{}, err
	}

	return updatedAlbum, nil
}

