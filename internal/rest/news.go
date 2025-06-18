package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"news-service/internal/newsportal"
	"strconv"
)

const (
	newsId = `id`
)

type ReqQuery struct {
	NewsFilter
	PageSize int `query:"page_size"`
	PageNum  int `query:"page_num"`
}

func (h *Handler) getOneNews(c echo.Context) error {
	ctx := c.Request().Context()
	newsIdStr := c.Param(newsId)
	newsId, err := strconv.Atoi(newsIdStr)
	if err != nil {
		return newErrorResponse(c, http.StatusBadRequest, `invalid id param, should be int`)
	}

	news, err := h.services.NewsByID(ctx, newsId)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, serverError)
	}

	if news == nil {
		return newErrorResponse(c, http.StatusNotFound, `news not found`)
	}

	return c.JSON(http.StatusOK, news)
}

func (h *Handler) getNews(c echo.Context) error {
	ctx := c.Request().Context()
	req := &ReqQuery{}
	if err := c.Bind(req); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, `invalid query param(s)`)
	}

	news, err := h.services.NewsByFilters(ctx,
		&newsportal.NewsFilter{
			CategoryID: req.CategoryID,
			TagID:      req.TagID,
		},
		req.PageNum, req.PageSize)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, serverError)
	} else if len(news) == 0 {
		return newErrorResponse(c, http.StatusNotFound, `news not found`)
	}

	res := make([]News, 0, len(news))
	for _, n := range news {
		res = append(res, newNews(n))
	}

	return c.JSON(http.StatusOK, res)
}

func (h *Handler) getNewsCount(c echo.Context) error {
	ctx := c.Request().Context()
	filter := new(newsportal.NewsFilter)
	if err := c.Bind(filter); err != nil {
		return newErrorResponse(c, http.StatusBadRequest, `invalid query param(s)`)
	}

	count, err := h.services.NewsCount(ctx, filter)
	if err != nil {
		return newErrorResponse(c, http.StatusInternalServerError, serverError)
	}

	return c.JSON(http.StatusOK, count)
}
