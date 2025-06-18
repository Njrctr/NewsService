package app

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
	"news-service/internal/db"
	"news-service/internal/newsportal"
	"news-service/internal/rest"
	"time"
)

type App struct {
	Cfg    *Config
	DB     *pg.DB
	server *rest.Handler
}

func New(cfg *Config, dbconn *pg.DB) *App {
	repository := db.NewRepository(dbconn)
	services := newsportal.New(repository)
	handlers := rest.New(services)

	return &App{
		Cfg:    cfg,
		DB:     dbconn,
		server: handlers,
	}
}

type Config struct {
	HttpPort int
	Env      string

	DB *pg.Options
}

func (a *App) Run(ctx context.Context) error {

	srv := a.server.Init()

	go func() {
		<-ctx.Done()
		stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(stopCtx); err != nil {
			log.Fatal(err)
		}
	}()

	return srv.Start(fmt.Sprintf(":%d", a.Cfg.HttpPort))

}
