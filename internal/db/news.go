package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func (r *Repository) NewsByID(ctx context.Context, id int) (*News, error) {

	news := &News{ID: id}
	query := r.db.ModelContext(ctx, news).
		Relation("Category").WherePK()

	query = prepareFilters(query, nil)
	err := query.Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("ошибка запроса: %w", err)
	}

	return news, nil
}

func (r *Repository) NewsByFilters(ctx context.Context, filter *NewsFilter, offset, limit int) ([]News, error) {

	var news []News
	query := r.db.ModelContext(ctx, &news).
		Relation("Category")

	query = prepareFilters(query, filter)

	err := query.Order(`publishedAt DESC`).Offset(offset).Limit(limit).Select()
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %w", err)
	}

	return news, nil
}

func (r *Repository) NewsCount(ctx context.Context, filter *NewsFilter) (int, error) {

	query := r.db.ModelContext(ctx, &News{})

	query = prepareFilters(query, filter)

	count, err := query.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func prepareFilters(query *orm.Query, filter *NewsFilter) *orm.Query {
	query = query.Where(`"news"."statusId" = ?`, 1).
		Where(`"category"."statusId" = ?`, 1).
		Where(`"news"."publishedAt" <= NOW()`)

	if filter != nil {
		if filter.CategoryID != 0 {
			query = query.Where(`"news"."categoryId" = ?`, filter.CategoryID)
		}
		if filter.TagID != 0 {
			query = query.Where(`? = ANY ("tagIds")`, filter.TagID)
		}
	}

	return query
}
