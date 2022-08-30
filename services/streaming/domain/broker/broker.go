package broker

import "context"

type Broker interface {
	Produce(topic string, b []byte) error
	Consume(ctx context.Context, topic string, consumeNewest bool) (dataChannel chan []byte, errorChannel chan error, closeConsumer func() error)
	DeleteTopic(topic string) error
}
