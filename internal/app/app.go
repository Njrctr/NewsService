package app

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"news-service/internal/db"
	"news-service/internal/newsportal"
	"news-service/internal/rest"
	"time"
)

type App struct {
	Cfg    *Config
	DB     *db.DbEngine
	server *rest.Handler
}

func New(cfg *Config, dbconn *db.DbEngine) *App {
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
	// http config
	HttpPort       int    `toml:"httpPort"`
	LogQueries     bool   `toml:"logQueries"`
	GinReleaseMode bool   `toml:"ginReleaseMode"`
	Env            string `toml:"env"`

	DB *db.DBConfig `toml:"db"`
}

func (c *Config) Validate() error {
	if c.HttpPort == 0 {
		return errors.New(`httpPort is zero`)
	}
	if err := c.DB.Validate(); err != nil {
		return fmt.Errorf(`failed to validate db config: %s`, err.Error())
	}
	return nil
}

func (a *App) Run(ctx context.Context) error {

	srv := &http.Server{
		Addr:        fmt.Sprintf(":%d", a.Cfg.HttpPort),
		Handler:     a.server.Init(),
		BaseContext: func(net.Listener) context.Context { return ctx },
	}
	go func() {
		<-ctx.Done()
		stopCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(stopCtx); err != nil {
			log.Fatal(err)
		}
	}()

	return srv.ListenAndServe()

}
