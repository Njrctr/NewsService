package newsportal

import (
	"context"
)

func (s *Service) Tags(ctx context.Context) ([]*Tag, error) {
	tags, err := s.repo.Tags(ctx)
	if err != nil {
		return nil, err
	}

	req := make([]*Tag, 0, len(tags))
	for _, tag := range tags {
		req = append(req, newTag(tag))
	}
	return req, nil
}
