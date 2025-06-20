package db

import (
	"github.com/go-pg/pg/v10"
)

type NewsFilter struct {
	CategoryID int
	TagID      int
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
