package repository

import (
	"context"
	"news-service/internal/db"
	"news-service/internal/structs"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
)

type NewsRepo interface {
	GetNewsByID(ctx context.Context, id int) (*structs.News, error)
	GetNews(ctx context.Context, filter *structs.NewsFilter, offset, limit uint) ([]*structs.News, error)
	GetNewsCount(ctx context.Context, filter *structs.NewsFilter) (int, error)
}

type TagRepo interface {
	GetTagsByIds(ctx context.Context, ids []int) ([]*structs.Tag, error)
}

type Repository struct {
	NewsRepo
	TagRepo
}

var Q = goqu.Dialect(`postgres`)

func NewRepository(db db.IDB) *Repository {
	return &Repository{
		NewsRepo: NewNewsRepository(db),
		TagRepo:  NewTagRepository(db),
	}
}
