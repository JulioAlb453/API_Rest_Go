package application

type EventPublisher interface {
	PublishEvent(eventType string, data map[string]interface{}) error
}