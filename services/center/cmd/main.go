package main

import (
	"avatar/pkg/base_repository"
	"avatar/services/center/config"
	"avatar/services/center/handler"
	neo4jUtil "avatar/services/center/infra/neo4j"
	"avatar/services/center/infra/neo4j/repository"
	"fmt"
	"log"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"google.golang.org/grpc"

	pb "avatar/services/center/protos"
)

type App struct {
	config      *config.Environment
	neo4jDriver neo4j.Driver
}

func main() {
	cfg := loadEnvironment()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.CenterServicePort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	neo4jDriver := neo4jUtil.ConnectNeo4j(cfg)
	baseRepository := base_repository.CreateBaseRepository(neo4jDriver)
	serviceTemplateRepository := repository.NewServiceTemplateRepository(*baseRepository)
	routineCategoryRepository := repository.NewRoutineCategoryRepository(*baseRepository)
	routineRepository := repository.NewroutineRepository(*baseRepository)
	serviceTemplateCategoryRepository := repository.NewServiceTemplateCategoryRepository(*baseRepository)
	avatarRepository := repository.NewAvatarRepository(*baseRepository)
	app := &App{
		config:      cfg,
		neo4jDriver: neo4jDriver,
	}
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcRecovery.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcRecovery.StreamServerInterceptor(),
		)),
	)
	server := handler.CreateServer(
		app.config,
		app.neo4jDriver,
		serviceTemplateRepository,
		routineCategoryRepository,
		routineRepository,
		serviceTemplateCategoryRepository,
		avatarRepository,
	)
	pb.RegisterCenterServer(grpcServer, server)
	log.Printf("[INFO] start http server listening %v", cfg.CenterServicePort)
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
