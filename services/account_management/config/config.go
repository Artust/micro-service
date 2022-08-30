package config

import (
	"errors"

	env "github.com/Netflix/go-env"
)

var (
	ErrInvalidEnv = errors.New("invalid env")
)

type Environment struct {
	AccountManagementPort      int    `env:"ACCOUNT_MANAGEMENT_PORT, required=true"`
	Neo4jUri                   string `env:"NEO4J_URI, required=true"`
	Neo4jUserName              string `env:"NEO4J_USERNAME, required=true"`
	Neo4jPassword              string `env:"NEO4J_PASSWORD, required=true"`
	S3URI                      string `env:"S3_URI, required=true"`
	JwtSecretKey               string `env:"JWT_SECRET_KEY, required=true"`
	FrontendUri                string `env:"FRONTEND_URI, required=true"`
	JwtExpirationHour          int    `env:"JWT_EXPIRATION_HOUR, required=true"`
	ResetTokenExpirationMinute int    `env:"RESET_TOKEN_EXPIRATION_MINUTE, required=true"`
	SendgridApiKey             string `env:"SENDGRID_API_KEY, required=true"`
	MailFrom                   string `env:"MAIL_FROM"`
}

func Load() (*Environment, error) {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		return nil, err
	}
	return &environment, nil
}
