package application

import (
	"API_ejemplo/src/album/domain"
	"context"
	"encoding/json"
	"log"
)

type ConsumeAlbumEventUseCase struct {
	repo domain.IAlbums
}

func NewConsumeAlbumEventUseCase(repo domain.IAlbums) *ConsumeAlbumEventUseCase {
	return &ConsumeAlbumEventUseCase{
		repo: repo,
	}
}

func (uc *ConsumeAlbumEventUseCase) Handle(msg []byte) error {
	log.Println("🎧 Procesando mensaje en caso de uso...")

	var album domain.Album
	err := json.Unmarshal(msg, &album)
	if err != nil {
		log.Printf("❌ Error al deserializar el mensaje: %v", err)
		return err
	}

	if album.Title == "" || album.Artist == "" {
		log.Println("❌ Datos del álbum incompletos")
		return nil
	}

	ctx := context.Background()
	err = uc.repo.Save(ctx, album)
	if err != nil {
		log.Printf("❌ Error al guardar el álbum: %v", err)
		return err
	}

	log.Printf("✅ Álbum guardado correctamente: %+v", album)
	return nil
}
