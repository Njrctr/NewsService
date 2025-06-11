package service

import (
	"context"
	"fmt"
	repo "news-service/internal/repository"
	"news-service/internal/structs"
	"news-service/internal/tools"
)

func GetNewsByID(ctx context.Context, id int) (*structs.News, error) {
	return repo.GetNewsByID(ctx, id)
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

	fmt.Println("tagMap:", tagMap)
	fmt.Println("tagIds:", tagIds)

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
