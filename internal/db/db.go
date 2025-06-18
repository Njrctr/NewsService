package db

import (
	"context"
	"github.com/go-pg/pg/v10"
)

func New(ctx context.Context, cfg *pg.Options) (*pg.DB, error) { // TODO Убрать в Мэйн
	db := pg.Connect(cfg)

	if err := db.Ping(ctx); err != nil {
		return nil, err
	}

	return db, nil
}

func TestDBCfg() *pg.Options {
	cfg := &pg.Options{
		Addr:     "localhost:5432",
		User:     "newsuser",
		Password: "akgj123cguygecuw3riu1y23",
		Database: "news-db",
	}

	return cfg
}
