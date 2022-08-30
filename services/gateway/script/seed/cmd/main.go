package main

import (
	"avatar/services/gateway/script/seed/config"
	neo4jDefault "avatar/services/gateway/script/seed/infra/neo4j"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/joho/godotenv"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	s3Session       *session.Session
	config          *config.Environment
	neo4j           neo4j.Driver
	urlImageDefault string
	urlSound        string
	urlAnimationKey string
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
	sess, err := session.NewSession(&aws.Config{
		Credentials:      credentials.NewStaticCredentials(cfg.S3Config.S3ID, cfg.S3Config.Secret, ""),
		S3ForcePathStyle: aws.Bool(true),
		Region:           aws.String(cfg.S3Config.RegionID),
		Endpoint:         aws.String(cfg.S3Config.Endpoint)},
	)
	if err != nil {
		return nil, err
	}
	return sess, nil
}
func main() {
	log.Println("starting seed...")
	cfg := loadEnvironment()
	neo4j := neo4jDefault.ConnectNeo4j(cfg)
	session, err := configS3(cfg)
	if err != nil {
		log.Fatalf("Failed to connect s3: %v", err)

	}
	urlImageDefault := "defaultImageRoutine.png"
	urlSound := "soundRoutine.wav"
	urlAnimationKey := "animationKeyRoutine.dat"
	app := &Server{
		s3Session:       session,
		config:          cfg,
		neo4j:           neo4j,
		urlImageDefault: urlImageDefault,
		urlSound:        urlSound,
		urlAnimationKey: urlAnimationKey,
	}
	app.CreateDefaultData()
}
func (c Server) CreateDefaultData() (string, error) {
	session := c.neo4j.NewSession(neo4j.SessionConfig{})
	defer session.Close()
	//create routine center
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		err := c.SaveToS3(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateAccountActivityHistory(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateAvatar(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateServiceTemplate(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateCenter(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateCorporation(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateDevice(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateIpCamera(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateMonitor(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateNote(tx)
		if err != nil {
			return "", err
		}
		err = c.CreatePermission(tx)
		if err != nil {
			return "", err
		}
		err = c.CreatePos(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateRoutineCategory(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateRoutine(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateShop(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateTalkSessionHistory(tx)
		if err != nil {
			return "", err
		}
		err = c.CreateTalkSession(tx)
		if err != nil {
			return "", err
		}
		return "", nil
	})
	if err != nil {
		log.Error("error when write transaction, error: ", err)
		return "", err
	}
	return "", nil
}
