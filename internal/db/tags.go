package db

import (
	"context"
	"fmt"
	"news-service/internal/db/tables"

	"github.com/jackc/pgx/v5"
)

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
