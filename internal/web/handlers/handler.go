package handlers

import (
	"context"
	"news-service/internal/structs"

	"github.com/gin-gonic/gin"
)

type NewsService interface {
	GetNewsByID(ctx context.Context, id int) (*structs.News, error)
	GetNews(ctx context.Context, filter *structs.NewsFilter, pageNum, pageSuze uint) ([]*structs.News, error)
	GetNewsCount(ctx context.Context, filter *structs.NewsFilter) (int, error)
}

type TagService interface {
}

type CategoryService interface {
}

type Service interface {
	NewsService
	TagService
	CategoryService
}

type Handler struct {
	news NewsService
	tag  TagService
	cat  CategoryService
}

func NewHandler(services Service) *Handler {
	return &Handler{
		news: services,
		tag:  services,
		cat:  services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {

	router := gin.Default()

	news := router.Group("/news")
	{
		news.GET("/", h.getNews)           // Получить новости по фильтру
		news.GET("/:id", h.getOneNews)     // Получить конкретную новость по id
		news.GET("/count", h.getNewsCount) // Получить количество новостей по фильтру
	}

	return router
}
