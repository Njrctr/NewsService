package newsportal

import (
	"context"
)

func (s *Service) GetCategories(ctx context.Context) ([]*Category, error) {
	cats, err := s.repo.GetCategories(ctx)
	if err != nil {
		return nil, err
	}
	req := make([]*Category, 0, len(cats))
	for _, cat := range cats {
		req = append(req, newCategory(cat))
	}
	return req, nil
}
