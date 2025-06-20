package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/go-pg/pg/v10"
	"github.com/k0kubun/pp"
	"log"
	"news-service/internal/app"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//go:generate zenrpc
func main() {
	ctx, shutdown := context.WithCancel(context.Background())

	var configFile string
	flag.StringVar(&configFile, `c`, `internal/cfg/.local.toml`, `config file name (*.toml)`)
	flag.Parse()

	cfg, err := loadConfig(configFile)
	if err != nil {
		log.Fatal(err)
	}

	dbconn := pg.Connect(cfg.DB)
	if err := dbconn.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	a := app.New(cfg, dbconn)
	a.InitServer()
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

func loadConfig(configFile string) (*app.Config, error) {
	cfg := &app.Config{}
	_, err := toml.DecodeFile(configFile, cfg)
	if err != nil {
		return nil, err
	}

	fmt.Printf("current config: \n\n%s\n", pp.Sprint(cfg))

	return cfg, nil
}
