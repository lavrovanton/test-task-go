package config

import (
	"log"
	"sync"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Port       string `env:"PORT"`
	PGHost     string `env:"POSTGRES_HOST"`
	PGPort     string `env:"POSTGRES_PORT"`
	PGDatabase string `env:"POSTGRES_DB"`
	PGUser     string `env:"POSTGRES_USER"`
	PGPassword string `env:"POSTGRES_PASSWORD"`
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		if err := env.Parse(&cfg); err != nil {
			log.Fatal(err)
		}
	})
	return &cfg
}
