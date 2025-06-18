package newsportal

import (
	"context"
	"news-service/internal/db"
)

func (s *Service) NewsByID(ctx context.Context, id int) (*News, error) {
	news, err := s.repo.NewsByID(ctx, id)
	if err != nil {
		return nil, err
	} else if news == nil {
		return nil, nil
	}

	req := newNews(news)

	if len(news.TagIDs) == 0 {
		return req, nil
	}
	tags, err := s.repo.Tags(ctx, news.TagIDs)
	if err != nil {
		return nil, err
	} else if len(tags) == 0 {
		return req, nil
	}

	for _, tag := range tags {
		req.Tags = append(req.Tags, newTag(tag))
	}

	return req, nil
}

func (s *Service) NewsByFilters(ctx context.Context, filter *NewsFilter, pageNum, pageSize int) ([]News, error) {

	offset, limit := pagination(pageNum, pageSize)
	news, err := s.repo.NewsByFilters(ctx, &db.NewsFilter{CategoryID: filter.CategoryID, TagID: filter.TagID}, offset, limit)
	if err != nil {
		return nil, err
	} else if len(news) == 0 {
		return nil, nil
	}

	req := make([]News, 0, len(news))
	uniqTags := make(map[int]Tag)
	var tagIds []int
	for _, n := range news {
		for _, tagID := range n.TagIDs {
			if _, ok := uniqTags[tagID]; !ok {
				tagIds = append(tagIds, tagID)
				uniqTags[tagID] = Tag{}
			}
		}
	}

	tags, err := s.repo.Tags(ctx, tagIds)
	if err != nil {
		return nil, err
	} else if len(tags) == 0 {
		return req, nil
	}

	for _, t := range tags {
		uniqTags[t.ID] = newTag(t)
	}

	for _, n := range news {
		newNewsItem := newNews(&n)
		for _, tagID := range n.TagIDs {
			if v, ok := uniqTags[tagID]; ok {
				newNewsItem.Tags = append(newNewsItem.Tags, v)
			}
		}
		req = append(req, *newNewsItem)
	}
	return req, nil
}

func (s *Service) NewsCount(ctx context.Context, filter *NewsFilter) (int, error) {
	count, err := s.repo.NewsCount(ctx, &db.NewsFilter{CategoryID: filter.CategoryID, TagID: filter.TagID})
	if err != nil {
		return 0, err
	}
	return count, nil
}
