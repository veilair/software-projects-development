package esSerializer

import "github.com/mehdihadeli/store-golang-microservice-sample/pkg/core/domain"

type EventSerializer interface {
	Serialize(event domain.IDomainEvent) (*EventSerializationResult, error)
	SerializeObject(data interface{}) (*EventSerializationResult, error)
	Deserialize(data []byte, eventType string, contentType string) (domain.IDomainEvent, error)
	DeserializeObject(data []byte, eventType string, contentType string) (interface{}, error)
	ContentType() string
}

type EventSerializationResult struct {
	Data        []byte
	ContentType string
	EventType   string
}
