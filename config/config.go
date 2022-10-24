package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

type Config struct {
	PgUser     string
	PgPassword string
	PgAddr     string
	PgDatabase string
}

var (
	cfg  Config
	once sync.Once
)

func Get() *Config {
	once.Do(func() {
		loadEnv()
		cfg = Config{
			PgUser:     os.Getenv(pgUser),
			PgPassword: os.Getenv(pgPassword),
			PgAddr:     os.Getenv(pgAddr),
			PgDatabase: os.Getenv(pgDatabase),
		}
	})
	return &cfg
}

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		err = godotenv.Load("./../.env")
		if err != nil {
			log.Fatal(err)
		}
	}
}
