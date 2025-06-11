package service

import (
	"context"
	repo "news-service/internal/repository"
	"news-service/internal/structs"
	"news-service/internal/tools"
)

func GetNewsByID(ctx context.Context, id int) (*structs.News, error) {
	news, err := repo.GetNewsByID(ctx, id)
	if err != nil {
		return nil, err
	}

	tags, err := repo.GetTagsByIds(ctx, news.TagIDs)
	if err != nil {
		return nil, err
	}

	for _, tag := range news.TagIDs {
		news.Tags = append(news.Tags, tags[tag])
	}

	return news, nil
}
func GetNews(ctx context.Context, filter *structs.NewsFilter, pageNum, pageSize uint) ([]*structs.News, error) {

	offset, limit := tools.Pagination(pageNum, pageSize)
	news, err := repo.GetNews(ctx, filter, offset, limit)
	if err != nil {
		return nil, err
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

	tags, err := repo.GetTagsByIds(ctx, tagIds)
	if err != nil {
		return nil, err
	}

	for _, n := range news {
		for _, tag := range n.TagIDs {
			n.Tags = append(n.Tags, tags[tag])
		}
	}

	return news, nil
}

func GetNewsCount(ctx context.Context, filter *structs.NewsFilter) (int, error) {
	return repo.GetNewsCount(ctx, filter)
}
