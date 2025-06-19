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

var (
	errNotFound    = echo.NewHTTPError(http.StatusNotFound, `not found`)
	errBadRequest  = echo.NewHTTPError(http.StatusBadRequest, `invalid query param(s)`)
	errServerError = echo.NewHTTPError(http.StatusInternalServerError, `server error`)
)

type ReqQuery struct {
	NewsFilter
	PageSize int `query:"page_size"`
	PageNum  int `query:"page_num"`
}

func (h *Handler) getByID(c echo.Context) error {
	ctx := c.Request().Context()
	newsIdStr := c.Param(newsId)
	newsId, err := strconv.Atoi(newsIdStr)
	if err != nil {
		return errBadRequest
	}

	news, err := h.services.NewsByID(ctx, newsId)
	if err != nil {
		return errServerError
	}

	if news == nil {
		return errNotFound
	}

	req := newNews(*news)

	return c.JSON(http.StatusOK, req)
}

func (h *Handler) getNews(c echo.Context) error {
	ctx := c.Request().Context()
	res := &ReqQuery{}
	if err := c.Bind(res); err != nil {
		return errNotFound
	}

	news, err := h.services.NewsByFilters(ctx,
		&newsportal.NewsFilter{
			CategoryID: res.CategoryID,
			TagID:      res.TagID,
		},
		res.PageNum, res.PageSize)
	if err != nil {
		return errServerError
	} else if len(news) == 0 {
		return errNotFound
	}

	req := newNewsSlice(news)

	return c.JSON(http.StatusOK, req)
}

func (h *Handler) getCount(c echo.Context) error {
	ctx := c.Request().Context()
	filter := new(newsportal.NewsFilter)
	if err := c.Bind(filter); err != nil {
		return errBadRequest
	}

	count, err := h.services.NewsCount(ctx, filter)
	if err != nil {
		return errServerError
	}

	return c.JSON(http.StatusOK, count)
}
