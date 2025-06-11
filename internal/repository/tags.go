package repository

import (
	"context"
	"fmt"
	d "news-service/internal/db"
	"news-service/internal/db/tables"
	"news-service/internal/structs"
)

func GetTagsByIds(ctx context.Context, ids []int) (map[int]*structs.Tag, error) {
	t := tables.Tags

	query, args, _ := Q.Select(
		t.ID,
		t.Title,
		t.StatusID,
	).From(t.T).Where(t.ID.In(ids)).ToSQL()

	rows, err := d.DB.Query(ctx, `get_tags`, query, args...)
	if err != nil {
		return nil, fmt.Errorf("get_tags return error: %w", err)
	}
	defer rows.Close()

	resMap := make(map[int]*structs.Tag)
	for rows.Next() {
		tagItem := new(structs.Tag)

		err := rows.Scan(
			&tagItem.ID,
			&tagItem.Title,
			&tagItem.StatusID,
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
