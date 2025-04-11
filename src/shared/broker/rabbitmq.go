package broker

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQBroker struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}


func NewRabbitMQBroker(uri string) (*RabbitMQBroker, error) {
	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %v", err)
	}

	err = ch.ExchangeDeclare(
		"album.events", 
		"topic",       
		true,           
		false,        
		false,         
		false,         
		nil,          
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare exchange: %v", err)
	}

	return &RabbitMQBroker{
		conn: conn,
		ch:   ch,
	}, nil
}

func (b *RabbitMQBroker) Publish(routingKey string, payload interface{}) error {
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err)
	}

	return b.ch.Publish(
		"albums.events", 
		"album.data",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
			Timestamp:   time.Now(),
		},
	)
}

func (b *RabbitMQBroker) Consume(queueName, bindingKey string, handler func(msg []byte) error) error {
    q, err := b.ch.QueueDeclare(
        queueName,
        true,  
        false, 
        false, 
        false, 
        nil,
    )
    if err != nil {
        log.Printf("❌ Error declarando cola: %v", err)
        return err
    }

    err = b.ch.QueueBind(
        queueName,
        "album.data",
        "album.events",
        false,
        nil,
    )
    if err != nil {
        log.Printf("❌ Error en binding de cola: %v", err)
        return err
    }

    msgs, err := b.ch.Consume(
        q.Name,
        "",
        false, 
        false,
        false,
        false,
        nil,
    )
    if err != nil {
        log.Printf("❌ Error al iniciar consumo: %v", err)
        return err
    }

    go func() {
        log.Println("📡 Esperando mensajes en la cola:", q.Name)

        for d := range msgs {
            log.Println("📩 Mensaje recibido:", string(d.Body))

            if err := handler(d.Body); err != nil {
                log.Printf("❌ Error procesando mensaje: %v", err)
                continue
            }

            d.Ack(false)
            log.Println("✅ Mensaje procesado correctamente.")
        }
    }()

    return nil
}



func (b *RabbitMQBroker) Close() {
	if b.ch != nil {
		b.ch.Close()
	}
	if b.conn != nil {
		b.conn.Close()
	}
}