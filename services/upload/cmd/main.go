package main

import (
	"avatar/pkg/logger"
	"fmt"

	"avatar/services/upload/config"
	"avatar/services/upload/routers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"

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
	session, err := configS3(cfg)
	if err != nil {
		log.Fatalf("Failed to connect s3: %v", err)

	}
	app := &routers.RestServer{
		Config:    cfg,
		Engine:    engine,
		S3Session: session,
	}
	routers.InitRouter(app)
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
