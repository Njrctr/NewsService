package newsportal

import (
	"context"
	"news-service/internal/db"
	"time"
)

// The NewsByID return News with included slice of Tag
func (s *Manager) NewsByID(ctx context.Context, id int) (*News, error) {
	news, err := s.repo.NewsByID(ctx, id, s.repo.FullNews()) // ???
	if err != nil {
		return nil, err
	} else if news == nil {
		return nil, nil
	}

	req := newNews(news)

	if len(news.TagIDs) == 0 {
		return req, nil
	}
	tags, err := s.repo.TagsByFilters(ctx, &db.TagSearch{IDs: news.TagIDs}, db.PagerNoLimit)
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

// The NewsByFilters return slice of News by NewsFilter, pageNum, pageSize. Include slice Tag into news items
func (s *Manager) NewsByFilters(ctx context.Context, filter *NewsFilter, pageNum, pageSize int) ([]News, error) {

	//news, err := s.repo.NewsByFilters(ctx, &db.NewsFilter{CategoryID: filter.CategoryID, TagID: filter.TagID}, offset, limit)
	news, err := s.repo.NewsByFilters(ctx, filter.toDB(), db.NewPager(pageNum, pageSize), s.repo.FullNews())
	if err != nil {
		return nil, err
	} else if len(news) == 0 {
		return nil, nil
	}

	req := make([]News, 0, len(news))
	uniqTags := make(map[int]Tag)
	var tagIds []int
	for _, n := range news {
		newNewsItem := newNews(&n)
		for _, tagID := range n.TagIDs {
			if _, ok := uniqTags[tagID]; !ok {
				tagIds = append(tagIds, tagID)
				uniqTags[tagID] = Tag{}
			}
		}
		req = append(req, *newNewsItem)
	}
	if len(tagIds) == 0 {
		return req, nil
	}

	tags, err := s.repo.TagsByFilters(ctx, &db.TagSearch{IDs: tagIds}, db.PagerNoLimit)
	if err != nil {
		return nil, err
	} else if len(tags) == 0 {
		return req, nil
	}

	for _, t := range tags {
		uniqTags[t.ID] = newTag(t)
	}

	for i, n := range news {
		for _, tagID := range n.TagIDs {
			if v, ok := uniqTags[tagID]; ok {
				req[i].Tags = append(req[i].Tags, v)
			}
		}
	}
	return req, nil
}

func (s *Manager) NewsCount(ctx context.Context, filter *NewsFilter) (int, error) {
	count, err := s.repo.CountNews(ctx, filter.toDB())
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (nf NewsFilter) toDB() *db.NewsSearch {
	var categoryID, tagID *int
	if nf.CategoryID != 0 {
		categoryID = ptr(nf.CategoryID)
	}
	if nf.TagID != 0 {
		tagID = ptr(nf.TagID)
	}

	return &db.NewsSearch{
		CategoryID:    categoryID,
		TagID:         tagID,
		PublishedAtLe: ptr(time.Now()),
	}
}
