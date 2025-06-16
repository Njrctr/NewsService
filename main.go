package main

import (
	"context"
	"flag"
	"log"
	"news-service/internal/app"
	"os"
	"os/signal"
	"syscall"
	"time"

	"news-service/internal/db"
)

// @title News Service API
// @version 1.0
// @description API Server for News
func main() {
	ctx, shutdown := context.WithCancel(context.Background())

	var configFile string
	flag.StringVar(&configFile, `c`, `.local.yml`, `config file name (*.yml)`)
	flag.Parse()

	cfg, err := app.LoadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}
	dbconn, err := db.New(ctx, cfg.DB)
	if err != nil {
		log.Fatal(err)
	}

	a := app.New(cfg, dbconn)
	go func() {
		if err := a.Run(ctx); err != nil {
			log.Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Server is shutting down...")
	shutdown()
	time.Sleep(1 * time.Second)
	os.Exit(1)
}
