package config

import (
	"github.com/caarlos0/env/v7"
)

type DBConfig struct {
	Host     string `env:"DB_HOST"     envDefault:"localhost"`
	Port     string `env:"DB_PORT"     envDefault:"5432"`
	DBName   string `env:"DB_NAME"     envDefault:"sampledb"`
	User     string `env:"DB_USER"     envDefault:"admin"`
	Password string `env:"DB_PASSWORD" envDefault:"admin"`
}

func GetDBConfig() (DBConfig, error) {
	dbconfig := DBConfig{}
	if err := env.Parse(&dbconfig); err != nil {
		panic(err)
	}
	return dbconfig, nil
}
