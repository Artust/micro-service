package config

import (
	"errors"

	"github.com/Netflix/go-env"
)

var (
	ErrInvalidEnv = errors.New("invalid env")
)

type Environment struct {
	AllowOrigins string `env:"CORS_ALLOW_ORIGINS_UPLOAD, required=true"`
	UploadPort   string `env:"UPLOAD_PORT, required=true"`
	S3ID         string `env:"S3ID, required=true"`
	S3Secret     string `env:"SECRET, required=true"`
	S3RegionID   string `env:"REGIONID, required=true"`
	S3Endpoint   string `env:"ENDPOINT, required=true"`
	S3Uri        string `env:"S3_URI, required=true"`
}

func Load() (*Environment, error) {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		return nil, err
	}
	return &environment, nil
}
