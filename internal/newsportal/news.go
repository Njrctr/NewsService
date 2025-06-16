package newsportal

import (
	"context"
	"fmt"
	"news-service/internal/db"
)

func (s *Service) NewsByID(ctx context.Context, id int) (*News, error) {
	news, err := s.repo.NewsByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if news == nil {
		return nil, nil
	}

	tags, err := s.repo.TagsByIds(ctx, news.TagIDs)
	if err != nil {
		return nil, err
	}

	req := newNews(news)
	for _, tag := range news.TagIDs {
		req.Tags = append(req.Tags, newTag(tags[tag]))
	}
	fmt.Printf("%#v", news)
	return req, nil
}
func (s *Service) NewsByFilters(ctx context.Context, filter *NewsFilter, pageNum, pageSize uint) ([]*News, error) {

	offset, limit := pagination(pageNum, pageSize)
	news, err := s.repo.NewsByFilters(ctx, &db.NewsFilter{filter.CategoryID, filter.TagID}, offset, limit)
	if err != nil {
		return nil, err
	}

	if len(news) == 0 {
		return nil, nil
	}

	tagMap := make(map[int]struct{})
	var tagIds []int
	for _, n := range news {
		for _, tag := range n.TagIDs {
			if _, ok := tagMap[tag]; !ok {
				tagMap[tag] = struct{}{}
				tagIds = append(tagIds, tag)
			}
		}
	}

	tags, err := s.repo.TagsByIds(ctx, tagIds)
	if err != nil {
		return nil, err
	}

	req := make([]*News, 0, len(news))
	for _, n := range news {
		newNews := newNews(n)
		for _, tag := range n.TagIDs {
			newNews.Tags = append(newNews.Tags, newTag(tags[tag]))
		}
		req = append(req, newNews)
	}

	return req, nil
}

func (s *Service) NewsCount(ctx context.Context, filter *NewsFilter) (int, error) {
	count, err := s.repo.NewsCount(ctx, &db.NewsFilter{filter.CategoryID, filter.TagID})
	if err != nil {
		return 0, err
	}
	return count, nil
}
