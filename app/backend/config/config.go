package config

import (
	"encoding/json"
	"os"

	"github.com/caarlos0/env/v7"
)

type (
	Config struct {
		Env      string `env:"ENV"     envDefault:"dev"`
		Port     string `env:"PORT"     envDefault:"8080"`
		LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	}

	DBConfig struct {
		Host     string `env:"DB_HOST"     envDefault:"localhost"`
		Port     string `env:"DB_PORT"     envDefault:"5432"`
		DBName   string `env:"DB_NAME"     envDefault:"sampledb"`
		User     string `env:"DB_USER"     envDefault:"admin"`
		Password string `env:"DB_PASSWORD" envDefault:"admin"`
	}

	Secret struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func GetConfig() (Config, error) {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		panic(err)
	}
	return config, nil
}

func GetDBConfig() (*DBConfig, error) {
	dbconfig := &DBConfig{}
	err := env.Parse(dbconfig)
	if err != nil {
		panic(err)
	}

	// AWS Secret Managerで管理されている場合はuser名とパスワードを上書きする
	dbSecret := os.Getenv("DB_SECRET")
	if dbSecret != "" {
		secret := Secret{}
		err := json.Unmarshal([]byte(dbSecret), &secret)
		if err != nil {
			return nil, err
		}
		dbconfig.User = secret.Username
		dbconfig.Password = secret.Password
	}

	return dbconfig, nil
}
