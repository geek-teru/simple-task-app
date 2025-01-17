package config

import (
	"github.com/caarlos0/env/v7"
)

type Config struct {
	Env      string `env:"ENV"     envDefault:"dev"`
	Port     string `env:"PORT"     envDefault:"8080"`
	LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
}

type DBConfig struct {
	Host     string `env:"DB_HOST"     envDefault:"localhost"`
	Port     string `env:"DB_PORT"     envDefault:"5432"`
	DBName   string `env:"DB_NAME"     envDefault:"sampledb"`
	User     string `env:"DB_USER"     envDefault:"admin"`
	Password string `env:"DB_PASSWORD" envDefault:"admin"`
}

func GetConfig() (Config, error) {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}
	return config, nil
}

func GetDBConfig() (DBConfig, error) {
	dbconfig := DBConfig{}
	if err := env.Parse(&dbconfig); err != nil {
		panic(err)
	}
	return dbconfig, nil
}
