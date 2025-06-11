package repository

import (
	"news-service/internal/db"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

var Q = goqu.Dialect(`postgres`)

var DB *db.DbEngine
