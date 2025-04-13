package broker

import (
	"log"
	
	"github.com/gorilla/websocket"
)

type WSPublisherImpl struct {
	clients map[*websocket.Conn]bool
}

func NewWSPublisher() *WSPublisherImpl {
	return &WSPublisherImpl{
		clients: make(map[*websocket.Conn]bool),
	}
}

func (p *WSPublisherImpl) PublishWSMessage(topic string, message interface{}) error {
	log.Printf("Enviando mensaje WS: %s - %v", topic, message)
	return nil
}