package config

import (
	"errors"

	env "github.com/Netflix/go-env"
)

var (
	ErrInvalidEnv = errors.New("invalid env")
)

type Environment struct {
	DedaultServiceAddress string `env:"DEFAULT_SERVICE_URI" mapstructure:"DEFAULT_SERVICE_URI"`
	Neo4JConf             struct {
		NEO4J_URI string `env:"NEO4J_URI" mapstructure:"NEO4J_URI"`
		UserName  string `env:"NEO4J_USERNAME" mapstructure:"NEO4J_USERNAME"`
		Password  string `env:"NEO4J_PASSWORD" mapstructure:"NEO4J_PASSWORD"`
	} `mapstructure:",squash"`
	S3Config struct {
		S3ID     string `env:"S3ID" mapstructure:"S3ID"`
		Secret   string `env:"SECRET" mapstructure:"SECRET"`
		RegionID string `env:"REGIONID" mapstructure:"REGIONID"`
		Endpoint string `env:"ENDPOINT" mapstructure:"ENDPOINT"`
		S3URI    string `env:"S3_URI,required=true"`
	} `mapstructure:",squash"`
	UploadServiceUri string `env:"UPLOAD_SERVICE_URI, required=true"`
}

func Load() (*Environment, error) {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		return nil, err
	}
	return &environment, nil
}
