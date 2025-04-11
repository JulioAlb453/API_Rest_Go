package adapters

import (
	"API_ejemplo/src/album/application"
	"github.com/streadway/amqp"
	"log"
)

type AlbumConsumer struct {
	useCase *application.ConsumeAlbumEventUseCase
}

func NewAlbumConsumer(useCase *application.ConsumeAlbumEventUseCase) *AlbumConsumer {
	return &AlbumConsumer{useCase: useCase}
}

func (ac *AlbumConsumer) StartConsuming(ch *amqp.Channel, queueName string) error {
	msgs, err := ch.Consume(
		queueName,
		"",
		true,  // autoAck
		false, // exclusive
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	go func() {
		for msg := range msgs {
			log.Println("📩 Mensaje recibido")
			err := ac.useCase.Handle(msg.Body)
			if err != nil {
				log.Printf("❌ Error procesando mensaje: %v", err)
			}
		}
	}()

	log.Println("🚀 Escuchando mensajes...")
	return nil
}
