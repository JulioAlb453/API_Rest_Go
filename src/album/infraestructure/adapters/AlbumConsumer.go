package adapters

import (
	"API_ejemplo/src/album/application"
	"API_ejemplo/src/album/domain"
	"context"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

type AlbumConsumer struct {
	useCase *application.UpdateAlbumsUseCase
}

func NewAlbumConsumer(useCase *application.UpdateAlbumsUseCase) *AlbumConsumer {
	return &AlbumConsumer{useCase: useCase}
}

func (ac *AlbumConsumer) StartConsuming(ch *amqp.Channel, queueName string) error {
	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			var album domain.Album
			err := json.Unmarshal(msg.Body, &album)
			if err != nil {
				log.Printf("❌ Error deserializando el mensaje: %v", err)
				continue
			}

			updatedAlbum, err := ac.useCase.Execute(context.Background(), album)
			if err != nil {
				log.Printf("❌ Error procesando mensaje de actualización: %v", err)
				continue
			}

			log.Printf("✅ Álbum actualizado: %v", updatedAlbum)
		}
	}()

	log.Println("🚀 Escuchando mensajes...")
	return nil
}
