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
	"news-service/internal/repository"
	"news-service/internal/service"
	"news-service/internal/web"
	"news-service/internal/web/handlers"
)

func main() {
	ctx, shutdown := context.WithCancel(context.Background())

	var configFile string
	flag.StringVar(&configFile, `c`, `.local.yml`, `config file name (*.yml)`)
	flag.Parse()

	g.Cfg = new(configs.Config)

	if err := configs.LoadConfig(configFile, g.Cfg); err != nil {
		log.Fatal(err)
	}

	db, err := db.NewDB(ctx)
	if err != nil {
		log.Fatal(err)
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handlers.NewHandler(services)

	go web.Run(ctx, handlers.InitRoutes())

	quit := make(chan os.Signal, 2)
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
	shutdown()
	time.Sleep(1 * time.Second)
	os.Exit(1)
}
