package repository

import (
	"context"
	"errors"
	"news-service/internal/db"
	"news-service/internal/db/dbtools"
	"news-service/internal/db/tables"
	myErrors "news-service/internal/errors"
	"news-service/internal/structs"

	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jackc/pgx/v5"
)

type NewsRepository struct {
	db db.IDB
}

func NewNewsRepository(db db.IDB) *NewsRepository {
	return &NewsRepository{db: db}
}

func (r *NewsRepository) GetNewsByID(ctx context.Context, id int) (*structs.News, error) {
	n := tables.News

	query, args, _ := Q.Select(
		n.ID,
		n.CategoryID,
		n.Title,
		n.Foreword,
		n.Content,
		n.TagIDs,
		n.CreatedAt,
		n.PublishedAt,
		n.StatusID,
	).From(n.T).Where(n.ID.Eq(id)).ToSQL()

	row, err := r.db.Query(ctx, `get_news_by_id`, query, args...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, myErrors.ErrNoRows
		}
		return nil, err
	}

	news, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByNameLax[structs.News])
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (r *NewsRepository) GetNews(ctx context.Context, filter *structs.NewsFilter, offset, limit uint) ([]*structs.News, error) {
	n := tables.News

	q := Q.Select(
		n.ID,
		n.CategoryID,
		n.Title,
		n.Foreword,
		n.Author,
		n.TagIDs,
		n.CreatedAt,
		n.PublishedAt,
	)

	var whereSet []exp.Expression
	if filter.CategoryID != nil && *filter.CategoryID != 0 {
		whereSet = append(whereSet, n.CategoryID.Eq(*filter.CategoryID))
	}
	if filter.TagID != nil && *filter.TagID != 0 {
		whereSet = append(whereSet, n.CategoryID.Eq(*filter.TagID))
	}

	if len(whereSet) != 0 {
		q.Where(whereSet...)
	}

	query, args, _ := q.Offset(offset).Limit(limit).ToSQL()
	rows, err := r.db.Query(ctx, `get_news`, query, args...)
	if err != nil {
		return nil, err
	}

	news, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[structs.News])
	if err != nil {
		return nil, err
	}

	return news, nil
}

func (r *NewsRepository) GetNewsCount(ctx context.Context, filter *structs.NewsFilter) (int, error) {
	n := tables.News

	q := Q.Select(dbtools.Count1()).From(n.T).Where()

	var whereSet []exp.Expression
	if filter.CategoryID != nil && *filter.CategoryID != 0 {
		whereSet = append(whereSet, n.CategoryID.Eq(*filter.CategoryID))
	}
	if filter.TagID != nil && *filter.TagID != 0 {
		whereSet = append(whereSet, n.CategoryID.Eq(*filter.TagID))
	}

	if len(whereSet) != 0 {
		q.Where(whereSet...)
	}

	query, args, _ := q.
		ToSQL()

	var count int
	if err := r.db.QueryRow(ctx, `get_news_count`, query, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}
