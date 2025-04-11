package application

import (
	"API_ejemplo/src/album/domain"
	"API_ejemplo/src/shared/broker"
	"context"
	"errors"
	"log"
	"time"
)

type UpdateAlbumsUseCase struct {
	repo         domain.IAlbums
	broker       *broker.RabbitMQBroker
	stockWarning int
}

func NewUpdateAlbumsUseCase(repo domain.IAlbums, broker *broker.RabbitMQBroker) *UpdateAlbumsUseCase {
	return &UpdateAlbumsUseCase{repo: repo,
		broker:       broker,
		stockWarning: 8,
	}
}

func (uc *UpdateAlbumsUseCase) Execute(ctx context.Context, album domain.Album) (domain.Album, error) {
    log.Printf("Received Stock: %d", album.Stock)

    // Verificación de campos obligatorios
    if album.Artist == "" || album.Title == "" {
        return domain.Album{}, errors.New("invalid input: Artista y Titulo son requeridos")
    }
	
    // Obtener el album existente
    existingAlbum, err := uc.repo.GetAlbumsById(ctx, album.Id)
    if err != nil {
        if errors.Is(err, domain.ErrAlbumNotFound) {
            return domain.Album{}, domain.ErrAlbumNotFound
        }
        return domain.Album{}, err
    }

    // Crear una variable auxiliar para almacenar el valor de stock anterior
    oldStock := existingAlbum.Stock

    // Mostrar log con los valores antiguos y nuevos
    log.Printf("Old Stock: %d, New Stock: %d", oldStock, album.Stock)
	
    // Actualizar el album
    existingAlbum.Title = album.Title
    existingAlbum.Artist = album.Artist
    existingAlbum.Year = album.Year
    existingAlbum.Stock = album.Stock  // Aquí se actualiza el stock
    existingAlbum.Price = album.Price
    existingAlbum.LastUpdated = time.Now()

    // Log para verificar si el stock de existingAlbum se actualiza correctamente
    log.Printf("Before publishing events, Stock is: %d", existingAlbum.Stock)
	
    // Actualizar el album en el repositorio
    updatedAlbum, err := uc.repo.Update(ctx, existingAlbum)
    if err != nil {
        return domain.Album{}, err
    }

    // Log después de la actualización
    log.Printf("Updated Stock: %d", updatedAlbum.Stock)

    // Preparar los datos del evento
    eventData := map[string]interface{}{
        "id": updatedAlbum.Id,
        "title": updatedAlbum.Title,
        "artist": updatedAlbum.Artist,
        "year": updatedAlbum.Year,
        "stock": updatedAlbum.Stock,
        "price": updatedAlbum.Price,
    }

    // Llamar a la función para publicar los eventos, pasando los valores de stock antes y después
    uc.publishEvents(eventData, oldStock, updatedAlbum)

    return updatedAlbum, nil
}


func (uc *UpdateAlbumsUseCase) publishEvents(eventData map[string]interface{}, oldStock int, updatedAlbum domain.Album) {
	log.Println("Publishing events...")

	log.Printf("Before comparison: Old Stock: %d, New Stock: %d", oldStock, updatedAlbum.Stock)

	if oldStock != updatedAlbum.Stock {
		log.Println("Stock has changed, publishing stock-related events...")

		stockEvent := make(map[string]interface{})
		for k, v := range eventData {
			stockEvent[k] = v
		}
		stockEvent["stock_change"] = updatedAlbum.Stock - oldStock

		if updatedAlbum.Stock < uc.stockWarning {
			log.Print("ento al if del stock")
			stockEvent["warning_level"] = "Baja cantidad"
			uc.publishEvent("album.stock.low", stockEvent)
		} else{
			log.Print("No entro al if")
		}

		if updatedAlbum.Stock <= 0 {
			stockEvent["warning_level"] = "Sin existencias"
			uc.publishEvent("album.stock.out", stockEvent)
		}
	} else {
		log.Println("Stock has not changed, no event will be published.")
	}
}




func (uc *UpdateAlbumsUseCase) publishEvent(eventType string, data map[string]interface{}) {
    event := map[string]interface{}{
        "event_type": eventType,
        "timestamp":  time.Now().UTC(),
        "data":       data,
    }

    err := uc.broker.Publish("album.events"+eventType, event)
    if err != nil {
        log.Printf("Failed to publish %s event: %v", eventType, err)
    }
}