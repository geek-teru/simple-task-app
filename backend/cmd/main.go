package main

import (
	"context"
	"log"
	"os"

	"github.com/geek-teru/simple-task-app/db"
	"github.com/geek-teru/simple-task-app/router"
)

func main() {
	// logger, err := log.New(cfg.LogLevel)
	// if err != nil {
	// 	log.Fatal("Failed to logger init:", err)
	// }
	// defer log.Sync(logger)

	// db setup
	client, err := db.NewClient()
	if err != nil {
		log.Fatal("Failed to DB Client init:", err)
	}
	defer client.Close()

	// Migration
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if err := client.Schema.Create(context.Background()); err != nil {
			log.Fatalf("Failed to creating schema resources: %v", err)
		}
	}

	// Server startup
	e := router.NewRouter()
	if err := e.Start(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
		os.Exit(1)
	}

}
