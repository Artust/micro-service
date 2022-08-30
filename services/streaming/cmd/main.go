package main

import (
	"avatar/services/streaming/config"
	"avatar/services/streaming/domain/broker"
	"context"
	"fmt"
	"log"
	"net"

	kafkaRepository "avatar/services/streaming/infra/kafka/repository"

	"avatar/services/streaming/handler"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"

	pb "avatar/services/streaming/protos"

	googleSpeech "cloud.google.com/go/speech/apiv1"
)

type App struct {
	config             *config.Environment
	brokerClient       broker.Broker
	googleSpeechClient *googleSpeech.Client
}

func main() {
	cfg := loadEnvironment()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.StreamingPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	kafkaClient := createBrokerClient(cfg)
	googleSpeechClient, err := createGoogleSpeechClient()
	if err != nil {
		log.Fatalf("Failed to create google speech client: %v", err)
	}
	app := &App{
		config:             cfg,
		brokerClient:       kafkaClient,
		googleSpeechClient: googleSpeechClient,
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcRecovery.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcRecovery.StreamServerInterceptor(),
		)),
	)
	server := handler.CreateServer(app.brokerClient, app.googleSpeechClient, app.config)
	pb.RegisterStreamingServer(grpcServer, server)
	log.Printf("[INFO] start http server listening %d", cfg.StreamingPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func loadEnvironment() *config.Environment {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("fail loading environment variables")
	}
	return cfg
}

func createBrokerClient(cfg *config.Environment) broker.Broker {
	client, err := kafkaRepository.Connect(cfg)
	if err != nil {
		log.Fatal("fail loading kafka:", err)
	}
	return client
}

func createGoogleSpeechClient() (*googleSpeech.Client, error) {
	client, err := googleSpeech.NewClient(context.Background())
	if err != nil {
		return nil, err
	}
	return client, nil
}
