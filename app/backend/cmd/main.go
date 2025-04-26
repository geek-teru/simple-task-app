package main

import (
	"context"
	"os"

	"github.com/geek-teru/simple-task-app/config"
	"github.com/geek-teru/simple-task-app/db"
	"github.com/geek-teru/simple-task-app/handler"
	"github.com/geek-teru/simple-task-app/log"
	"github.com/geek-teru/simple-task-app/repository"
	"github.com/geek-teru/simple-task-app/router"
	"github.com/geek-teru/simple-task-app/service"
	echo "github.com/labstack/echo/v4"
)

func main() {
	cfg, nil := config.GetConfig()

	logger, err := log.New(cfg.LogLevel)
	if err != nil {
		panic(err)
	}
	defer log.Sync(logger)

	// db setup
	client, err := db.NewClient()
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// Migration
	if len(os.Args) > 1 && os.Args[1] == "migrate" {
		if err := client.Schema.Create(context.Background()); err != nil {
			panic(err)
		}
	} else {
		// User
		userRepo := repository.NewUserRepository(client)
		userService := service.NewUserService(userRepo)
		userHandler := handler.NewUserHandler(userService, logger)

		// Task
		taskRepo := repository.NewTaskRepository(client)
		taskService := service.NewTaskService(taskRepo)
		taskHandler := handler.NewTaskHandler(taskService, logger)

		e := echo.New()
		// e.HideBanner = true
		router.NewRouter(e, *userHandler, *taskHandler)

		// Server startup
		if err := e.Start(":" + cfg.Port); err != nil {
			panic(err)
		}
	}
	os.Exit(0)
}
