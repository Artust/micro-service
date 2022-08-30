package config

import (
	"errors"

	env "github.com/Netflix/go-env"
)

var (
	ErrInvalidEnv = errors.New("invalid env")
)

type Environment struct {
	KafkaServerURI           string `env:"KAFKA_SERVERS_URI,required=true"`
	StreamingPort            int    `env:"STREAMING_PORT,required=true"`
	SpeechToTextLanguageCode string `env:"SPEECH_T0_TEXT_LANGUAGE_CODE,required=true"`
	SpeechToTextSampleRate   int    `env:"SPEECH_TO_TEXT_SAMPLE_RATE, required=true"`
}

func Load() (*Environment, error) {
	var environment Environment
	_, err := env.UnmarshalFromEnviron(&environment)
	if err != nil {
		return nil, err
	}
	return &environment, nil
}
