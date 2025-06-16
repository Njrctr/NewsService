package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/k0kubun/pp"
	"gopkg.in/yaml.v3"
	"log"
	"net"
	"net/http"
	"news-service/internal/db"
	"news-service/internal/newsportal"
	"news-service/internal/rest"
	"os"
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
	HttpPort       int    `yaml:"httpPort"`
	LogQueries     bool   `yaml:"logQueries"`
	GinReleaseMode bool   `yaml:"ginReleaseMode"`
	Env            string `yaml:"env"`

	DB *db.Config `yaml:"db"`
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

func LoadConfig(fileName string) (*Config, error) {
	file, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf(`cant read file '%s'`, fileName)
	}

	config := new(Config)
	if err = yaml.Unmarshal(file, config); err != nil {
		return nil, fmt.Errorf(`file %s yaml unmarshal error: %s`, fileName, err.Error())
	}
	if config.Env == "dev" {
		fmt.Printf("current config: \n\n%s\n", pp.Sprint(config))
	}
	return config, config.Validate()
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
