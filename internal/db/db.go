package db

import (
	"context"
	"errors"
	"github.com/go-pg/pg/v10"
	"strconv"
	"time"
)

type DBConfig struct {
	Host        string        `toml:"host"`
	Port        int           `toml:"port"`
	Username    string        `toml:"username"`
	Password    string        `toml:"password"`
	Database    string        `toml:"database"`
	MaxConn     int           `toml:"maxConn"`
	MinConn     int           `toml:"minConn"`
	MaxIdleTime time.Duration `toml:"maxIdleTime"`
	TimeZone    string        `toml:"timezone"`
	DisableTLS  bool          `toml:"disableTLS"`
}

func (c *DBConfig) Validate() error {
	if c.Host == `` {
		return errors.New(`empty host`)
	}
	if c.Port == 0 {
		return errors.New(`port is zero`)
	}
	if c.Username == `` {
		return errors.New(`empty username`)
	}
	if c.Password == `` {
		return errors.New(`empty password`)
	}
	if c.Database == `` {
		return errors.New(`empty database`)
	}
	if c.MaxConn == 0 {
		return errors.New(`maxCons is zero`)
	}
	if c.MaxIdleTime < time.Second {
		return errors.New(`maxIdleTime is less than 1 second`)
	}
	if c.TimeZone == `` {
		return errors.New(`empty timezone`)
	}

	return nil
}

func New(ctx context.Context, cfg *DBConfig) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     cfg.Host + ":" + strconv.Itoa(cfg.Port),
		User:     cfg.Username,
		Password: cfg.Password,
		Database: cfg.Database,
	})

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func TestDBCfg() *DBConfig {
	cfg := &DBConfig{
		Host:        "localhost",
		Port:        5432,
		Username:    "newsuser",
		Password:    "akgj123cguygecuw3riu1y23",
		Database:    "news-db",
		MaxConn:     300,
		MinConn:     10,
		MaxIdleTime: 10,
		TimeZone:    "Europe/Moscow",
		DisableTLS:  true,
	}

	return cfg
}
