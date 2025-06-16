package db

import (
	"context"
	"errors"
	"fmt"
	"news-service/internal/db/dbtools"
	"news-service/internal/db/tables"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) NewsByID(ctx context.Context, id int) (*News, error) {
	n := tables.News
	c := tables.Categories

	query, args, _ := Q.Select(
		n.ID,
		n.Title,
		n.Foreword,
		n.Content,
		n.Author,
		n.TagIDs,
		n.CreatedAt,
		n.PublishedAt,
		c.ID.As(`category_id`),
		c.Title.As(`category_title`),
	).From(n.T).Join(c.T, goqu.On(n.CategoryID.Eq(c.ID))).Where(n.ID.Eq(id), n.StatusID.Eq(1), n.PublishedAt.Lte(dbtools.Now())).ToSQL()

	row, err := r.db.Query(ctx, `get_news_by_id`, query, args...)
	if err != nil {
		return nil, fmt.Errorf("ошибка запроса: %w", err)
	}

	news, err := pgx.CollectOneRow(row, func(row pgx.CollectableRow) (*News, error) {
		var (
			categoryID    int
			categoryTitle string
			newsTitle     string
		)

		newsItem := &News{}

		err := row.Scan(
			&newsItem.ID,
			&newsTitle,
			&newsItem.Foreword,
			&newsItem.Content,
			&newsItem.Author,
			&newsItem.TagIDs,
			&newsItem.CreatedAt,
			&newsItem.PublishedAt,
			&categoryID,
			&categoryTitle,
		)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %w", err)
		}

		newsItem.Category = &Category{
			ID:    categoryID,
			Title: categoryTitle,
		}
		newsItem.Title = newsTitle

		return newsItem, nil
	})
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("collect failed: %w", err)
	}

	return news, nil
}

func (r *Repository) NewsByFilters(ctx context.Context, filter *NewsFilter, offset, limit uint) ([]*News, error) {
	n := tables.News
	c := tables.Categories

	q := Q.Select(
		n.ID,
		n.CategoryID,
		n.Title,
		n.Foreword,
		n.Author,
		n.TagIDs,
		n.CreatedAt,
		n.PublishedAt,
		c.ID.As(`category_id`),
		c.Title.As(`category_title`),
	).From(n.T).Where(n.StatusID.Eq(1), n.PublishedAt.Lte(dbtools.Now())).Join(c.T, goqu.On(n.CategoryID.Eq(c.ID), c.StatusID.Eq(1)))

	if filter != nil {
		includeFilter(q, filter)
	}

	query, args, _ := q.Order(n.PublishedAt.Desc()).Offset(offset).Limit(limit).ToSQL()

	rows, err := r.db.Query(ctx, `get_news`, query, args...)
	if err != nil {

		return nil, fmt.Errorf("coalesce error: %w", err)
	}

	news, err := pgx.CollectRows(rows, func(row pgx.CollectableRow) (*News, error) {
		var (
			categoryID    int
			categoryTitle string
			newsTitle     string
		)

		newsItem := &News{}

		err := row.Scan(
			&newsItem.ID,
			&newsItem.CategoryID,
			&newsTitle,
			&newsItem.Foreword,
			&newsItem.Author,
			&newsItem.TagIDs,
			&newsItem.CreatedAt,
			&newsItem.PublishedAt,
			&categoryID,
			&categoryTitle,
		)
		if err != nil {
			return nil, fmt.Errorf("ошибка сканирования строки: %w", err)
		}

		newsItem.Category = &Category{
			ID:    categoryID,
			Title: categoryTitle,
		}
		newsItem.Title = newsTitle

		return newsItem, nil
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка сбора строк: %w", err)
	}

	return news, nil
}

func (r *Repository) NewsCount(ctx context.Context, filter *NewsFilter) (int, error) {
	n := tables.News

	q := Q.Select(dbtools.Count1()).From(n.T).Where(n.StatusID.Eq(1), n.PublishedAt.Lte(dbtools.Now()))

	if filter != nil {
		includeFilter(q, filter)
	}

	query, args, _ := q.
		ToSQL()

	var count int
	if err := r.db.QueryRow(ctx, `get_news_count`, query, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func includeFilter(s *goqu.SelectDataset, filter *NewsFilter) {
	n := tables.News

	var whereSet []exp.Expression
	if filter.CategoryID != 0 {
		whereSet = append(whereSet, n.CategoryID.Eq(filter.CategoryID))
	}
	if filter.TagID != 0 {
		whereSet = append(whereSet, n.CategoryID.Eq(filter.TagID))
	}

	*s = *s.Where(whereSet...)
}
