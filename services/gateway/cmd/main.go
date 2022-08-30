package main

import (
	"avatar/pkg/logger"
	"fmt"
	"net"
	"sync"
	"time"

	"avatar/services/gateway/config"
	"avatar/services/gateway/domain/broker"
	handlerGrpc "avatar/services/gateway/handler/grpc"
	kafkaRepository "avatar/services/gateway/infra/kafka/repository"
	ulRepository "avatar/services/gateway/infra/upload/respository"
	pbAccountManagement "avatar/services/gateway/protos/account_management"
	pbCenter "avatar/services/gateway/protos/center"
	pbCorporation "avatar/services/gateway/protos/corporation"
	pb "avatar/services/gateway/protos/gateway"
	pbPos "avatar/services/gateway/protos/pos"
	pbStreaming "avatar/services/gateway/protos/streaming"
	pbTalkSession "avatar/services/gateway/protos/talk_session"
	"avatar/services/gateway/routers"
	"log"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
	cfg := loadEnvironment()
	engine := gin.New()
	// Logger
	if gin.IsDebugging() {
		engine.Use(gin.Logger())
		engine.Use(gin.Recovery())
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		engine.Use(logger.SetLogger())
		engine.Use(logger.Recovery())
	}
	accountManagementClient, err := createAccountManagementClient(cfg)
	if err != nil {
		log.Fatalf("Failed to connect pos service: %v", err)
	}
	posClient, err := createPosClient(cfg)
	if err != nil {
		log.Fatalf("Failed to connect pos service: %v", err)
	}
	centerClient, err := createCenterClient(cfg)
	if err != nil {
		log.Fatalf("Failed to connect center service: %v", err)
	}
	corporationClient, err := createCorporationClient(cfg)
	if err != nil {
		log.Fatalf("Failed to connect center service: %v", err)
	}
	streamingClient, err := createStreamingClient(cfg)
	if err != nil {
		log.Fatalf("Failed to connect center service: %v", err)
	}
	talkSessionClient, err := createTalkSessionClient(cfg)
	if err != nil {
		log.Fatalf("Failed to connect center service: %v", err)
	}
	session, err := configS3(cfg)
	if err != nil {
		log.Fatalf("Failed to connect s3: %v", err)
	}
	upload := ulRepository.NewUploadRepository(cfg.UploadServiceUri, config.RoutineBucketName)
	kafkaClient := createBrokerClient(cfg)
	app := &routers.RestServer{
		StreamingClient:         streamingClient,
		AccountManagementClient: accountManagementClient,
		PosClient:               posClient,
		CenterClient:            centerClient,
		TalkSessionClient:       talkSessionClient,
		CorporationClient:       corporationClient,
		BrokerClient:            kafkaClient,
		Config:                  cfg,
		Engine:                  engine,
		S3Session:               session,
		Upload:                  *upload,
	}
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		routers.InitRouter(app)
		wg.Done()
	}()
	// register grpc
	go func() {
		keepAliveConfig := keepalive.ServerParameters{
			Time:    10 * time.Second, // Ping the client if it is idle for 5 seconds to ensure the connection is still active
			Timeout: 5 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
		}
		log.Println("gateway port", cfg.GatewayGrpcPort)
		lis, err := net.Listen("tcp", fmt.Sprintf(":%v", cfg.GatewayGrpcPort))
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
			wg.Done()
		}
		grpcServer := grpc.NewServer(
			grpc.UnaryInterceptor(grpcMiddleware.ChainUnaryServer(
				grpcRecovery.UnaryServerInterceptor(),
			)),
			grpc.StreamInterceptor(grpcMiddleware.ChainStreamServer(
				grpcRecovery.StreamServerInterceptor(),
			)),
			grpc.KeepaliveParams(keepAliveConfig),
		)
		server := handlerGrpc.CreateServer(streamingClient, talkSessionClient)
		pb.RegisterAvatarServer(grpcServer, server)
		log.Printf("[INFO] start http server listening %d", cfg.GatewayGrpcPort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
			wg.Done()
		}
	}()
	wg.Wait()
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

func createAccountManagementClient(cfg *config.Environment) (pbAccountManagement.AccountManagementClient, error) {
	log.Println("handler service address: ", cfg.AccountManagementServiceURI)
	connection, err := grpc.Dial(cfg.AccountManagementServiceURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pbAccountManagement.NewAccountManagementClient(connection)
	return client, nil
}

func createCenterClient(cfg *config.Environment) (pbCenter.CenterClient, error) {
	log.Println("handler service address: ", cfg.CenterServiceURI)
	connection, err := grpc.Dial(cfg.CenterServiceURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pbCenter.NewCenterClient(connection)
	return client, nil
}

func createPosClient(cfg *config.Environment) (pbPos.POSClient, error) {
	log.Println("handler service address: ", cfg.POSServiceURI)
	connection, err := grpc.Dial(cfg.POSServiceURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pbPos.NewPOSClient(connection)
	return client, nil
}

func createCorporationClient(cfg *config.Environment) (pbCorporation.CorporationClient, error) {
	log.Println("handler service address: ", cfg.CorporationServiceUri)
	connection, err := grpc.Dial(cfg.CorporationServiceUri, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pbCorporation.NewCorporationClient(connection)
	return client, nil
}

func createStreamingClient(cfg *config.Environment) (pbStreaming.StreamingClient, error) {
	log.Println("handler service address: ", cfg.StreamingURI)
	connection, err := grpc.Dial(cfg.StreamingURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pbStreaming.NewStreamingClient(connection)
	return client, nil
}

func createTalkSessionClient(cfg *config.Environment) (pbTalkSession.TalkSessionClient, error) {
	log.Println("handler service address: ", cfg.TalkSessionServiceURI)
	connection, err := grpc.Dial(cfg.TalkSessionServiceURI, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pbTalkSession.NewTalkSessionClient(connection)
	return client, nil
}

func configS3(cfg *config.Environment) (*session.Session, error) {
	session, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(cfg.S3ID, cfg.S3Secret, ""),
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(cfg.S3RegionID),
		Endpoint:         aws.String(cfg.S3Endpoint)},
	)
	if err != nil {
		return nil, err
	}
	return session, nil

}

func createBrokerClient(cfg *config.Environment) broker.Broker {
	client, err := kafkaRepository.Connect(cfg)
	if err != nil {
		log.Fatal("fail loading kafka:", err)
	}
	return client
}
