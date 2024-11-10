package config

import (
	"fmt"
	"reflect"

	"github.com/caarlos0/env/v11"
)

type Settings struct {
	Address          string        `env:"ADDRESS,required,notEmpty"`
	Environnement    Environnement `env:"ENVIRONNEMENT,required,notEmpty"`
	SentryDSN        string        `env:"SENTRY_DSN"`
	PostgresPassword string        `env:"POSTGRES_PASSWORD,required,notEmpty"`
	PostgresUser     string        `env:"POSTGRES_PASSWORD,required,notEmpty"`
	PostgresDb       string        `env:"POSTGRES_PASSWORD,required,notEmpty"`
}

var envOptions = env.Options{
	FuncMap: map[reflect.Type]env.ParserFunc{
		reflect.TypeOf(Environnement{}): parseEnvironnement,
	},
}

func validateSettings(settings *Settings) error {
	if settings.Environnement.IsProd() && settings.SentryDSN == "" {
		return fmt.Errorf("sentry DSN is required when environment is prod")
	}
	return nil
}

func LoadSettings() (*Settings, error) {
	settings := &Settings{}
	if err := env.ParseWithOptions(settings, envOptions); err != nil {
		return nil, err
	}
	if err := validateSettings(settings); err != nil {
		return nil, err
	}
	return settings, nil
}
