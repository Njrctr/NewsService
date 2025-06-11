package repository

import (
	"context"
	"errors"
	"fmt"
	d "news-service/internal/db"
	"news-service/internal/db/dbtools"
	"news-service/internal/db/tables"
	"news-service/internal/structs"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jackc/pgx/v5"
)

func GetNewsByID(ctx context.Context, id int) (*structs.News, error) {
	n := tables.News
	c := tables.Categories

	query, args, _ := Q.Select(
		n.ID,
		c.ID,
		c.Title.As(`category_title`),
		c.StatusID.As(`category_status`),
		n.Title,
		n.Foreword,
		n.Content,
		n.TagIDs,
		n.CreatedAt,
		n.PublishedAt,
		n.StatusID,
	).From(n.T).Join(c.T, goqu.On(n.CategoryID.Eq(c.ID))).Where(n.ID.Eq(id)).ToSQL()

	row, err := d.DB.Query(ctx, `get_news_by_id`, query, args...)
	if err != nil {
		return nil, err
	}

	news, err := pgx.CollectOneRow(row, pgx.RowToAddrOfStructByNameLax[structs.News])
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("collect failed: %w", err)
	}

	return news, nil
}

func GetNews(ctx context.Context, filter *structs.NewsFilter, offset, limit uint) ([]*structs.News, error) {
	n := tables.News
	c := tables.Categories

	q := Q.Select(
		n.ID,
		n.CategoryID,
		c.ID.As(`category_id`),
		c.Title.As(`category_title`),
		c.StatusID.As(`category_status`),
		n.Title,
		n.StatusID,
		n.Foreword,
		n.Author,
		n.TagIDs,
		n.CreatedAt,
		n.PublishedAt,
	).From(n.T).Join(c.T, goqu.On(n.CategoryID.Eq(c.ID)))

	var whereSet []exp.Expression
	if filter.CategoryID != 0 {
		whereSet = append(whereSet, n.CategoryID.Eq(filter.CategoryID))
	}
	if filter.TagID != 0 {
		whereSet = append(whereSet, n.CategoryID.Eq(filter.TagID))
	}

	query, args, _ := q.Where(whereSet...).Offset(offset).Limit(limit).ToSQL()

	rows, err := d.DB.Query(ctx, `get_news`, query, args...)
	if err != nil {

		return nil, fmt.Errorf("coalesce error: %w", err)
	}

	news, err := pgx.CollectRows(rows, newsScanner)
	if err != nil {
		return nil, fmt.Errorf("ошибка сбора строк: %w", err)
	}

	return news, nil
}

func GetNewsCount(ctx context.Context, filter *structs.NewsFilter) (int, error) {
	n := tables.News

	q := Q.Select(dbtools.Count1()).From(n.T).Where()

	var whereSet []exp.Expression
	if filter.CategoryID != 0 {
		whereSet = append(whereSet, n.CategoryID.Eq(filter.CategoryID))
	}
	if filter.TagID != 0 {
		whereSet = append(whereSet, n.CategoryID.Eq(filter.TagID))
	}

	if len(whereSet) != 0 {
		q.Where(whereSet...)
	}

	query, args, _ := q.
		ToSQL()

	var count int
	if err := d.DB.QueryRow(ctx, `get_news_count`, query, args...).Scan(&count); err != nil {
		return 0, err
	}

	return count, nil
}

func newsScanner(row pgx.CollectableRow) (*structs.News, error) {
	var (
		categoryID       int
		categoryTitle    string
		categoryStatusID int
		newsTitle        string
		newsStatusID     int
	)

	newsItem := &structs.News{}

	err := row.Scan(
		&newsItem.ID,
		&newsItem.CategoryID,
		&categoryID,
		&categoryTitle,
		&categoryStatusID,
		&newsTitle,
		&newsStatusID,
		&newsItem.Foreword,
		&newsItem.Author,
		&newsItem.TagIDs,
		&newsItem.CreatedAt,
		&newsItem.PublishedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("ошибка сканирования строки: %w", err)
	}

	newsItem.Category = &structs.Category{
		ID:       categoryID,
		Title:    categoryTitle,
		StatusID: categoryStatusID,
	}
	newsItem.Title = newsTitle
	newsItem.StatusID = newsStatusID

	return newsItem, nil
}
