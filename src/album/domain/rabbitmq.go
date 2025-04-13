package domain

type RabbitMQ interface {
	Consume(queueName, bindingKey string, handler func(msg []byte)) error
}
 