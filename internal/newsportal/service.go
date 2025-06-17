package newsportal

import (
	"context"
	"news-service/internal/db"
)

type Service struct {
	repo *db.Repository
}

func New(repo *db.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Tags(ctx context.Context) ([]Tag, error) {
	tags, err := s.repo.Tags(ctx, nil)
	if err != nil {
		return nil, err
	}

	req := make([]Tag, 0, len(tags))
	for _, tag := range tags {
		req = append(req, newTag(tag))
	}
	return req, nil
}

func (s *Service) GetCategories(ctx context.Context) ([]Category, error) {
	cats, err := s.repo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	req := make([]Category, 0, len(cats))
	for _, cat := range cats {
		req = append(req, newCategory(cat))
	}
	return req, nil
}

func pagination(pageNum, pageSize int) (int, int) {
	if pageSize == 0 {
		pageSize = 5
	}

	page := pageNum
	if page > 0 {
		page--
	}

	return pageSize * page, pageSize
}
