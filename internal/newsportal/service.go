package newsportal

import (
	"context"
	"news-service/internal/db"
)

type Manager struct {
	repo db.NewsRepo
}

func NewManager(repo db.NewsRepo) *Manager {
	return &Manager{repo: repo.WithEnabledOnly()}
}

// The Tags return slice of Tag
func (s *Manager) Tags(ctx context.Context) ([]Tag, error) {
	tags, err := s.repo.TagsByFilters(ctx, nil, db.PagerNoLimit)
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
	cats, err := s.repo.CategoriesByFilters(ctx, nil, db.PagerNoLimit, db.WithSort(db.NewSortField(db.Columns.Category.OrderNumber, false)))
	if err != nil {
		return nil, err
	}
	req := make([]Category, 0, len(cats))
	for _, cat := range cats {
		req = append(req, newCategory(cat))
	}
	return req, nil
}

func ptr[T any](t T) *T {
	return &t
}
