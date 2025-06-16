package db

import (
	"context"
	"fmt"
	"news-service/internal/db/tables"

	"github.com/jackc/pgx/v5"
)

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
