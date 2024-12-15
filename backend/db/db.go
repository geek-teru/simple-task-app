package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/geek-teru/simple-task-app/config"
	"github.com/geek-teru/simple-task-app/ent"
)

func NewClient() (*ent.Client, error) {
	conf, err := config.GetDBConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	//defer client.Close()

	return client, nil
}

func NewDB() (*sql.DB, error) {
	conf, err := config.GetDBConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		conf.Host, conf.Port, conf.User, conf.Password, conf.DBName)

	//postgres
	db, err := sql.Open("postgres", dsn)

	if err != nil {
		panic(err)
	}
	return db, nil
}
