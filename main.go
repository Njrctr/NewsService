package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"news-service/internal/configs"
	"news-service/internal/db"
	g "news-service/internal/global"
	"news-service/internal/web"
	"news-service/internal/web/handlers"
)

// @title News Service API
// @version 1.0
// @description API Server for News
func main() {
	ctx, shutdown := context.WithCancel(context.Background())

	var configFile string
	flag.StringVar(&configFile, `c`, `.local.yml`, `config file name (*.yml)`)
	flag.Parse()

	cfg := new(configs.Config)

	if err := configs.LoadConfig(configFile, cfg); err != nil {
		log.Fatal(err)
	}

	if err := db.InitDB(ctx, cfg.DB); err != nil {
		log.Fatal(err)
	}

	handlers := handlers.InitRoutes()

	go web.Run(ctx, handlers, cfg.App)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

Loop:
	for {
		select {
		case err := <-g.Errors:
			log.Fatalf("fatal err:%s", err.Error())
		case <-quit:
			break Loop
		}
	}

	log.Print("Server is shutting down...")
	shutdown() // TODO Разобрать с GracefullShutdown
	time.Sleep(1 * time.Second)
	os.Exit(1)
}
