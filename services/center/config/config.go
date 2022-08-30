package config

import (
	"errors"

	env "github.com/Netflix/go-env"
)

var (
	ErrInvalidEnv = errors.New("invalid env")
)

type Environment struct {
	CenterServicePort int    `env:"CENTER_PORT, required=true"`
	POSServiceAddress string `env:"POS_SERVICE_URI, required=true"`
	Neo4jUri          string `env:"NEO4J_URI, required=true"`
	Neo4jUserName     string `env:"NEO4J_USERNAME, required=true"`
	Neo4jPassword     string `env:"NEO4J_PASSWORD, required=true"`

	S3ID     string `env:"S3ID, required=true"`
	Secret   string `env:"SECRET, required=true"`
	RegionID string `env:"REGIONID, required=true"`
	Endpoint string `env:"ENDPOINT, required=true"`
	S3URI    string `env:"S3_URI, required=true"`
}

func Load() (*Environment, error) {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		return nil, err
	}
	return &environment, nil
}
