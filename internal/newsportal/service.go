package newsportal

import (
	"context"
	"news-service/internal/db"
)

type Manager struct {
	repo *db.Repository
}

func NewManager(repo *db.Repository) *Manager {
	return &Manager{repo: repo}
}

// The Tags return slice of Tag
func (s *Manager) Tags(ctx context.Context) ([]Tag, error) {
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

// The Categories return slice of Category
func (s *Manager) Categories(ctx context.Context) ([]Category, error) {
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
