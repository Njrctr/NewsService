package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"news-service/internal/newsportal"
)

type Handler struct {
	services *newsportal.Service
}

func New(services *newsportal.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init() *echo.Echo {

	router := echo.New()

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

func (h *Handler) getTags(c echo.Context) error {
	ctx := c.Request().Context()
	tags, err := h.services.Tags(ctx)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, serverError)
	}

	return c.JSON(http.StatusOK, tags)
}

func (h *Handler) getCategories(c echo.Context) error {
	ctx := c.Request().Context()
	cats, err := h.services.GetCategories(ctx)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, serverError)
	}

	return c.JSON(http.StatusOK, cats)
}
