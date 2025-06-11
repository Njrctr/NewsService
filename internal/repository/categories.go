package repository

import (
	"context"
	"fmt"
	d "news-service/internal/db"
	"news-service/internal/db/tables"
	"news-service/internal/structs"

	"github.com/jackc/pgx/v5"
)

func GetCategories(ctx context.Context) ([]*structs.Category, error) {
	c := tables.Categories

	query, args, _ := Q.Select(
		c.ID,
		c.Title,
		c.StatusID,
		c.OrderNumber,
	).From(c.T).Order(c.OrderNumber.Asc()).ToSQL()

	rows, err := d.DB.Query(ctx, `get_categories`, query, args...)
	if err != nil {
		return nil, fmt.Errorf("get_categories query return err: %w", err)
	}

	cats, err := pgx.CollectRows(rows, pgx.RowToAddrOfStructByNameLax[structs.Category])
	if err != nil {
		return nil, fmt.Errorf("get_categories failed to collect rows: %w", err)
	}

	return cats, nil
}
