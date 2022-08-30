package repository

import (
	"context"
	"fmt"

	"github.com/Shopify/sarama"
)

func (c *Client) Consume(ctx context.Context, topic string, consumeNewest bool) (
	dataChannel chan []byte,
	errorChannel chan error,
	closeConsumer func() error,
) {
	consumer := c.consumer
	messagesChannel := make(chan []byte)
	errorsChannel := make(chan error)
	partitions, _ := consumer.Partitions(topic)
	var offset int64
	if consumeNewest {
		offset = sarama.OffsetNewest
	} else {
		offset = sarama.OffsetOldest
	}
	messages, err := consumer.ConsumePartition(topic, partitions[0], offset)
	if err != nil {
		fmt.Println(err)
		return
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
