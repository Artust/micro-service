package main

import (
	"avatar/pkg/base_repository"
	"avatar/services/account_management/config"
	"avatar/services/account_management/handler"
	neo4jUtil "avatar/services/account_management/infra/neo4j"
	"avatar/services/account_management/infra/neo4j/repository"
	"avatar/services/account_management/infra/sendgrid"
	pb "avatar/services/account_management/protos"
	"fmt"
	"log"
	"net"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	cfg := loadEnvironment()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.AccountManagementPort))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	neo4jDriver := neo4jUtil.ConnectNeo4j(cfg)
	baseRepository := base_repository.CreateBaseRepository(neo4jDriver)
	accountRepository := repository.NewAccountRepository(*baseRepository)
	permissionRepository := repository.NewPermissionRepository(*baseRepository)
	accountRoleRepository := repository.NewAccountRoleRepository(*baseRepository)
	resetPasswordTokenRepository := repository.NewResetPasswordTokenRepository(*baseRepository)
	userActivityRepository := repository.NewUserActivityRepository(*baseRepository)
	mailClient := sendgrid.NewMailClient(cfg)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
			grpcRecovery.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
			grpcRecovery.StreamServerInterceptor(),
		)),
	)
	server := handler.CreateServer(
		accountRepository,
		permissionRepository,
		accountRoleRepository,
		resetPasswordTokenRepository,
		userActivityRepository,
		mailClient,
		cfg,
		neo4jDriver)
	pb.RegisterAccountManagementServer(grpcServer, server)
	log.Printf("[INFO] start http server listening %d", cfg.AccountManagementPort)
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
		log.Fatal("fail loading environment variables, error: ", err)
	}
	return cfg
}
