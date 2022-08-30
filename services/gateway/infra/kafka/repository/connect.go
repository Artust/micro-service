package repository

import (
	"avatar/services/gateway/config"
	"log"
	"strings"

	"github.com/Shopify/sarama"
)

type Client struct {
	producer     sarama.SyncProducer
	consumer     sarama.Consumer
	clusterAdmin sarama.ClusterAdmin
}

func Connect(cfg *config.Environment) (*Client, error) {
	producer, err := createProducer(cfg)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}
	consumer, err := createConsumer(cfg)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	clusterAdmin, err := createClusterAdmin(cfg)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	client := &Client{
		producer:     producer,
		consumer:     consumer,
		clusterAdmin: clusterAdmin,
	}

	return client, nil
}

func createProducer(cfg *config.Environment) (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(strings.Split(cfg.KafkaServerURI, ","), config)
	if err != nil {
		return nil, err
	}
	return producer, nil
}

func createConsumer(cfg *config.Environment) (sarama.Consumer, error) {
	config := sarama.NewConfig()
	consumer, err := sarama.NewConsumer(strings.Split(cfg.KafkaServerURI, ","), config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

func createClusterAdmin(cfg *config.Environment) (sarama.ClusterAdmin, error) {
	config := sarama.NewConfig()
	clusterAdmin, err := sarama.NewClusterAdmin(strings.Split(cfg.KafkaServerURI, ","), config)
	if err != nil {
		return nil, err
	}
	return clusterAdmin, nil
}
