package repository

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
)

func (c *Client) Consume(ctx context.Context, topic string) (
	dataChannel chan []byte,
	errorChannel chan error,
	closeConsumer func() error,
) {
	consumer := c.consumer
	messagesChannel := make(chan []byte)
	errorsChannel := make(chan error)
	partitions, _ := consumer.Partitions(topic)
	messages, err := consumer.ConsumePartition(topic, partitions[0], sarama.OffsetOldest)
	if err != nil {
		fmt.Println(err)
	}
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case err := <-messages.Errors():
				if err != nil {
					errorsChannel <- err
					fmt.Println("consumerError: ", err)
				}
			case msg := <-messages.Messages():
				if msg != nil {
					messagesChannel <- msg.Value
				}
			}
		}
	}()
	closeConsumer = func() error {
		return messages.Close()
	}
	return messagesChannel, errorsChannel, closeConsumer
}
