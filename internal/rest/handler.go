package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news-service/internal/newsportal"
)

type Handler struct {
	services *newsportal.Service
}

func New(services *newsportal.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init() *gin.Engine {

	router := gin.Default()

	news := router.Group("/news")
	{
		news.GET("/", h.getNews)           // Получить новости по фильтру
		news.GET("/:id", h.getOneNews)     // Получить конкретную новость по id
		news.GET("/count", h.getNewsCount) // Получить количество новостей по фильтру
	}

	tags := router.Group("/tags")
	{
		tags.GET("/", h.getTags) // Получить все теги
	}

	categories := router.Group("/categories")
	{
		categories.GET("/", h.getCategories) // Получить все категории
	}

	return router
}

func (h *Handler) getTags(c *gin.Context) {
	tags, err := h.services.Tags(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, serverError)
		return
	}

	c.JSON(http.StatusOK, tags)
}

func (h *Handler) getCategories(c *gin.Context) {
	cats, err := h.services.GetCategories(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, serverError)
		return
	}

	c.JSON(http.StatusOK, cats)
}
