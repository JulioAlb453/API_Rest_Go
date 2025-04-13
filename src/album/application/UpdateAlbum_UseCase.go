// application/use_case.go
package application

import (
	"API_ejemplo/src/album/domain"
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"
)

type UpdateAlbumsUseCase struct {
	repo          domain.IAlbums
	rabbitMQBroker RabbitMQPublisher
	broadcaster   domain.Broadcaster
	stockWarning  int
}

type RabbitMQPublisher interface {
	Publish(queue string, message []byte) error
}

func NewUpdateAlbumsUseCase(repo domain.IAlbums, rabbitMQBroker RabbitMQPublisher, broadcaster domain.Broadcaster) *UpdateAlbumsUseCase {
	return &UpdateAlbumsUseCase{
		repo:          repo,
		broadcaster:   broadcaster,
		rabbitMQBroker: rabbitMQBroker,
		stockWarning:  8,
	}
}

func (uc *UpdateAlbumsUseCase) Execute(ctx context.Context, album domain.Album) (domain.Album, error) {
	if album.Artist == "" || album.Title == "" {
		return domain.Album{}, errors.New("artista y título son requeridos")
	}
	if album.Price <= 0 || album.Stock < 0 || album.Year <= "" {
		return domain.Album{}, errors.New("precio, stock y año deben ser válidos")
	}

	existingAlbum, err := uc.repo.GetAlbumsById(ctx, album.Id)
	if err != nil {
		return domain.Album{}, err
	}

	existingAlbum.Title = album.Title
	existingAlbum.Artist = album.Artist
	existingAlbum.Year = album.Year
	existingAlbum.Stock = album.Stock
	existingAlbum.Price = album.Price
	existingAlbum.LastUpdated = time.Now()

	updatedAlbum, err := uc.repo.Update(ctx, existingAlbum)
	if err != nil {
		return domain.Album{}, err
	}

	if updatedAlbum.Stock < uc.stockWarning {
		alertMessage := map[string]interface{}{
			"event_type": "low_stock_alert",
			"id":         updatedAlbum.Id,
			"title":      updatedAlbum.Title,
			"stock":      updatedAlbum.Stock,
			"threshold":  uc.stockWarning,
			"timestamp":  time.Now().Format(time.RFC3339),
		}

		if jsonAlert, err := json.Marshal(alertMessage); err == nil {
			if err := uc.rabbitMQBroker.Publish("stock_alerts", jsonAlert); err != nil {
				log.Printf("❌ Error enviando alerta de stock bajo: %v", err)
			} else {
				log.Printf("🚨 Alerta enviada: low_stock_alert -> %s", jsonAlert)
			}

			log.Println("🔔 Enviando alerta al WebSocket:", string(jsonAlert))
			uc.broadcaster.BroadcastMessage(jsonAlert)
			log.Print("Alerta enviada al WS")
		} else {
			log.Printf("❌ Error al crear el mensaje de alerta: %v", err)
		}
	}

	return updatedAlbum, nil
}
