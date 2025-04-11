package broker

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type RabbitMQBroker struct{
	conn *amqp.Connection
	ch 	 *amqp.Channel
	retryCount int
	retryDelay time.Duration
	exchangeName string
}

func NewRabbiMQBroker(uri string) (*RabbitMQBroker, error){
	config := &RabbitMQBroker{
		retryCount: 3,
		retryDelay: 2 *time.Second,
		exchangeName: "album_events",
	}
	var err error
	for i := 0; i < config.retryCount; i++ {
		config.conn, err = amqp.Dial(uri)
		if err == nil {
			break
		}
		log.Printf("Failed to connect to RabbitMQ (attempt %d/%d): %v", i+1, config.retryCount, err)
		time.Sleep(config.retryDelay)
	}
	config.ch, err = config.conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("failed to open channel: %v", err)
	}

	err = config.ch.ExchangeDeclare(
		config.exchangeName,
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

	return config, nil
}

func (b *RabbitMQBroker) Publish(topic string, payload interface{}) error{
	body, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error marshaling payload: %v", err)
	}

	return b.ch.Publish(
		b.exchangeName,
		topic,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body: body,
			Timestamp: time.Now(),
		},
	)
}

func (b *RabbitMQBroker) Consume(queueName, topic string) (<-chan amqp.Delivery, error) {
	q, err := b.ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		amqp.Table{
			"x-dead-letter-exchange":    "dlx",
			"x-dead-letter-routing-key": "dlx.routing.key",
		},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue: %v", err)
	}

	err = b.ch.QueueBind(
		q.Name,
		topic,
		b.exchangeName,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to bind queue: %v", err)
	}

	return b.ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
}

func (b *RabbitMQBroker) Close() {
	if b.ch != nil {
		b.ch.Close()
	}
	if b.conn != nil {
		b.conn.Close()
	}
}