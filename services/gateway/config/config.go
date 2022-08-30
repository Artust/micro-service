package config

import (
	"errors"

	"github.com/Netflix/go-env"
)

var (
	ErrInvalidEnv = errors.New("invalid env")
)

type Environment struct {
	GatewayGrpcPort             int    `env:"GATEWAY_GRPC_PORT, required=true"`
	GatewayRestPort             int    `env:"GATEWAY_REST_PORT, required=true"`
	AllowOrigins                string `env:"CORS_ALLOW_ORIGINS_GATEWAY, required=true"`
	S3ID                        string `env:"S3ID, required=true"`
	S3Secret                    string `env:"SECRET, required=true"`
	S3RegionID                  string `env:"REGIONID, required=true"`
	S3Endpoint                  string `env:"ENDPOINT, required=true"`
	S3Uri                       string `env:"S3_URI, required=true"`
	Neo4jUri                    string `env:"NEO4J_URI, required=true"`
	Neo4jUserName               string `env:"NEO4J_USERNAME, required=true"`
	Neo4jPassword               string `env:"NEO4J_PASSWORD, required=true"`
	AccountManagementServiceURI string `env:"ACCOUNT_MANAGEMENT_SERVICE_URI, required=true"`
	POSServiceURI               string `env:"POS_SERVICE_URI, required=true"`
	CenterServiceURI            string `env:"CENTER_SERVICE_URI, required=true"`
	TalkSessionServiceURI       string `env:"TALK_SESSION_SERVICE_URI, required=true"`
	StreamingURI                string `env:"STREAMING_URI, required=true"`
	KafkaServerURI              string `env:"KAFKA_SERVERS_URI, required=true"`
	CorporationServiceUri       string `env:"CORPORATION_SERVICE_URI, required=true"`
	UploadServiceUri            string `env:"UPLOAD_SERVICE_URI, required=true"`
	JwtSecretKey                string `env:"JWT_SECRET_KEY, required=true"`
	JwtExpirationHour           int64  `env:"JWT_EXPIRATION_HOUR, required=true"`
	ServerRtspUri               string `env:"SERVER_RTSP_URI, required=true"`
}

func Load() (*Environment, error) {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		return nil, err
	}
	return &environment, nil
}
