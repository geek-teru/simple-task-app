package main

import (
	"context"
	"log"

	"github.com/geek-teru/simple-task-app/db"
)

func main() {
	client, err := db.NewClient()
	if err != nil {
		panic(err)
	}

	// Migration
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}