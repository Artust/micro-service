package repository

import (
	"time"

	"github.com/Shopify/sarama"
)

func (c *Client) Produce(topic string, b []byte) error {
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(b),
		Key:   sarama.StringEncoder(time.Now().String()),
	}
	if _, _, err := c.producer.SendMessage(msg); err != nil {
		return err
	}
	return nil
}
