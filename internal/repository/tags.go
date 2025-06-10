package repository

import (
	"context"
	"news-service/internal/db"
	"news-service/internal/structs"
)

type TagRepository struct {
	db db.IDB
}

func NewTagRepository(db db.IDB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) GetTagsByIds(ctx context.Context, ids []int) ([]*structs.Tag, error) {
	return nil, nil
}
