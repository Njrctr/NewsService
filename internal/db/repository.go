package db

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	"github.com/jackc/pgx/v5"
	"news-service/internal/db/tables"
)

var Q = goqu.Dialect(`postgres`)

type Repository struct {
	db *DbEngine
}

func NewRepository(db *DbEngine) *Repository {
	return &Repository{db: db}
}

func (r *Repository) TagsByIds(ctx context.Context, ids []int) (map[int]*Tag, error) {
	t := tables.Tags

	query, args, _ := Q.Select(
		t.ID,
		t.Title,
	).From(t.T).Where(t.ID.In(ids)).ToSQL()

	rows, err := r.db.Query(ctx, `get_tags_by_ids`, query, args...)
	if err != nil {
		return nil, fmt.Errorf("get_tags return error: %w", err)
	}
	defer rows.Close()

	resMap := make(map[int]*Tag)
	for rows.Next() {
		tagItem := new(Tag)

		err := rows.Scan(
			&tagItem.ID,
			&tagItem.Title,
		)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %w", err)
		}

		resMap[tagItem.ID] = tagItem
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("ошибка итерации по строкам: %w", err)
	}

	return resMap, nil
}

func (r *Repository) Tags(ctx context.Context) ([]*Tag, error) {
	t := tables.Tags

	query, args, _ := Q.Select(
		t.ID,
		t.Title,
	).From(t.T).Where(t.StatusID.Eq(1)).ToSQL()

	rows, err := r.db.Query(ctx, `get_tags`, query, args...)
	if err != nil {
		return nil, fmt.Errorf("get_tags query return err: %w", err)
	}

	tags, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[Tag])
	if err != nil {
		return nil, fmt.Errorf("get_tags failed to collect rows: %w", err)
	}

	return tags, nil
}

func (r *Repository) GetCategories(ctx context.Context) ([]*Category, error) {
	c := tables.Categories

	query, args, _ := Q.Select(
		c.ID,
		c.Title,
		c.OrderNumber,
	).From(c.T).Order(c.OrderNumber.Asc()).ToSQL()

	rows, err := r.db.Query(ctx, `get_categories`, query, args...)
	if err != nil {
		return nil, fmt.Errorf("get_categories query return err: %w", err)
	}

	cats, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[Category])
	if err != nil {
		return nil, fmt.Errorf("get_categories failed to collect rows: %w", err)
	}

	return cats, nil
}
