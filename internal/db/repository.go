package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-pg/pg/v10"
)

type Repository struct {
	db *pg.DB
}

func NewRepository(db *pg.DB) *Repository {
	return &Repository{db: db}
}

type NewsFilter struct {
	CategoryID int
	TagID      int
}

func (r *Repository) Tags(ctx context.Context, ids []int) ([]Tag, error) {

	var tags []Tag
	query := r.db.ModelContext(ctx, &tags).
		Where(`? = ?`, pg.Ident(Columns.Tag.StatusID), 1)

	if len(ids) != 0 {
		where := fmt.Sprintf("%q IN (?)", Columns.Tag.ID)
		query = query.WhereIn(where, ids)
	}
	err := query.Select()
	if err != nil {
		if errors.Is(err, pg.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("get_tags failed: %w", err)
	}

	return tags, nil
}

func (r *Repository) GetCategories(ctx context.Context) ([]Category, error) {

	var cats []Category
	err := r.db.ModelContext(ctx, &cats).
		Where(`? = ?`, pg.Ident(Columns.Category.StatusID), 1).
		Order(fmt.Sprintf("%s ASC", pg.Ident(Columns.Category.OrderNumber))).Select()
	if err != nil {
		return nil, fmt.Errorf("get_categories query return err: %w", err)
	}

	return cats, nil
}

func TestDBCfg() *pg.Options {
	cfg := &pg.Options{
		Addr:     "localhost:5432",
		User:     "newsuser",
		Password: "akgj123cguygecuw3riu1y23",
		Database: "news-db",
	}

	return cfg
}
