package broker

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
	"github.com/streadway/amqp"
)

type RabbitMQBroker struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func NewRabbitMQBroker(url string) (*RabbitMQBroker, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatalf("Error al conectar a RabbitMQ: %v", err)
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		log.Fatalf("Error al abrir un canal de RabbitMQ: %v", err)
		return nil, err
	}

	return &RabbitMQBroker{
		conn: conn,
		ch:   ch,
	}, nil
}

func (b *RabbitMQBroker) PublishEvent(eventType string, data map[string]interface{}) error {
	event := map[string]interface{}{
		"event_type": eventType,
		"timestamp":  time.Now().UTC(),
		"data":       data,
	}

	body, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("error marshaling event: %v", err)
	}

	err = b.ch.Publish(
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

	if err != nil {
		log.Printf("Error al publicar evento %s: %v", eventType, err)
	}
	return err
}

func (rb *RabbitMQBroker) Consume(queueName, bindingKey string, handler func(msg []byte)) error {
	ch, err := rb.conn.Channel()
	if err != nil {
		return err
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
		return err
	}

	_, err = ch.QueueDeclare(
		queueName,
		true,      
		false,    
		false,     
		false,     
		nil,       
	)
	if err != nil {
		return err
	}

	err = ch.QueueBind(
		queueName,  
		bindingKey, 
		"album.events", 
		false,      
		nil,         
	)
	if err != nil {
		return err
	}

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
		for d := range msgs {
			handler(d.Body) 
		}
	}()

	return nil
}

func (b *RabbitMQBroker) Publish(queue string, message []byte) error {
	ch, err := b.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	_, err = ch.QueueDeclare(
		queue, 
		true,  
		false, 
		false, 
		false, 
		nil,   
	)
	if err != nil {
		return err
	}

	return ch.Publish(
		"",   
		queue,
		false, 
		false, 
		amqp.Publishing{
			ContentType: "application/json",
			Body:       message,
		})
}

func (b *RabbitMQBroker) Close() {
	if b.ch != nil {
		b.ch.Close()
	}
	if b.conn != nil {
		b.conn.Close()
	}
}
