package db

import (
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

var Q = goqu.Dialect(`postgres`)

type Repository struct {
	db *DbEngine
}

func NewRepository(db *DbEngine) *Repository {
	return &Repository{db: db}
}
