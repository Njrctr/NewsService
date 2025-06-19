package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"strings"
)

func (r *Repository) NewsByID(ctx context.Context, id int) (*News, error) {

	news := &News{ID: id}
	query := r.db.ModelContext(ctx, news).
		Relation(Columns.News.Category).WherePK()

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
		Relation(Columns.News.Category)

	query = prepareFilters(query, filter)

	order := fmt.Sprintf("%s DESC", pg.Ident(Columns.News.PublishedAt))

	err := query.Order(order).Offset(offset).Limit(limit).Select()
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %w", err)
	}

	return news, nil
}

func (r *Repository) NewsCount(ctx context.Context, filter *NewsFilter) (int, error) {

	query := r.db.ModelContext(ctx, &News{}).Relation(Columns.News.Category)

	query = prepareFilters(query, filter)

	count, err := query.Count()
	if err != nil {
		return 0, err
	}

	return count, nil
}

func prepareFilters(query *orm.Query, filter *NewsFilter) *orm.Query {
	categoryAlias := strings.ToLower(Columns.News.Category)
	newsAlias := Tables.News.Alias

	query = query.Where(`?.? = ?`, pg.Ident(newsAlias), pg.Ident(Columns.News.StatusID), 1).
		Where(`?.? = ?`, pg.Ident(categoryAlias), pg.Ident(Columns.Category.StatusID), 1).
		Where(`?.? <= NOW()`, pg.Ident(newsAlias), pg.Ident(Columns.News.PublishedAt))

	if filter != nil {
		if filter.CategoryID != 0 {
			query = query.Where(`?.? = ?`, pg.Ident(newsAlias), pg.Ident(Columns.News.CategoryID), filter.CategoryID)
		}
		if filter.TagID != 0 {
			query = query.Where(`? = ANY (?.?)`, filter.TagID, pg.Ident(newsAlias), pg.Ident(Columns.News.TagIDs))
		}
	}

	return query
}
