package service

import (
	"context"
	"news-service/internal/structs"
)

type TagRepo interface {
	GetTagsByIds(ctx context.Context, ids []int) ([]*structs.Tag, error)
}
type TagServiceI interface {
	GetTagsByIds(ctx context.Context, ids []int) ([]*structs.Tag, error)
}

type TagService struct {
	repo TagRepo
}

func NewTagService(repo TagRepo) *TagService {
	return &TagService{repo: repo}
}

func (s *TagService) GetTagsByIds(ctx context.Context, ids []int) ([]*structs.Tag, error) {
	return s.repo.GetTagsByIds(ctx, ids)
}
