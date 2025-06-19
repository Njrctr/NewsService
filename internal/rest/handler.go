package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"news-service/internal/newsportal"
)

type Handler struct {
	services *newsportal.Manager
}

func New(services *newsportal.Manager) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init() *echo.Echo {

	router := echo.New()

	news := router.Group("/news")
	{
		news.GET("/", h.getNews)       // Получить новости по фильтру
		news.GET("/:id", h.getByID)    // Получить конкретную новость по id
		news.GET("/count", h.getCount) // Получить количество новостей по фильтру
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
		return errServerError
	}

	req := newTagsSlice(tags)

	return c.JSON(http.StatusOK, req)
}

func (h *Handler) getCategories(c echo.Context) error {
	ctx := c.Request().Context()
	cats, err := h.services.Categories(ctx)
	if err != nil {
		return errServerError
	}

	req := newCategorySlice(cats)

	return c.JSON(http.StatusOK, req)
}
