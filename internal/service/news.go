package service

import (
	"context"
	"news-service/internal/structs"
	"news-service/internal/tools"
)

type NewsRepo interface {
	GetNewsByID(ctx context.Context, id int) (*structs.News, error)
	GetNews(ctx context.Context, filter *structs.NewsFilter, offset, limit uint) ([]*structs.News, error)
	GetNewsCount(ctx context.Context, filter *structs.NewsFilter) (int, error)
}

type NewsServiceI interface {
	GetNewsByID(ctx context.Context, id int) (*structs.News, error)
	GetNews(ctx context.Context, filter *structs.NewsFilter, pageNum, pageSize uint) ([]*structs.News, error)
	GetNewsCount(ctx context.Context, filter *structs.NewsFilter) (int, error)
}

type NewsService struct {
	newsRepo NewsRepo
	tagRepo  TagRepo
}

func NewNewsService(newsRepo NewsRepo, tagRepo TagRepo) *NewsService {
	return &NewsService{
		newsRepo: newsRepo,
		tagRepo:  tagRepo,
	}
}

func (s *NewsService) GetNewsByID(ctx context.Context, id int) (*structs.News, error) {
	return s.newsRepo.GetNewsByID(ctx, id)
}
func (s *NewsService) GetNews(ctx context.Context, filter *structs.NewsFilter, pageNum, pageSize uint) ([]*structs.News, error) {
	offset, limit := tools.Pagination(pageNum, pageSize)
	return s.newsRepo.GetNews(ctx, filter, offset, limit)
}

func (s *NewsService) GetNewsCount(ctx context.Context, filter *structs.NewsFilter) (int, error) {
	return s.newsRepo.GetNewsCount(ctx, filter)
}
