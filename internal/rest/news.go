package rest

import (
	"net/http"
	"news-service/internal/newsportal"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	newsId = `id`
)

type ReqQuery struct {
	NewsFilter
	PageSize int `form:"page_size,default=5"`
	PageNum  int `form:"page_num,default=0"`
}

// @Summary Получить новость по ID
// @Tags News
// @Description Get News By ID
// @ID get-news-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} structs.News
// @Failure 400,404,500 {object} errorResponse
// @Router /news [get]
func (h *Handler) getOneNews(c *gin.Context) {
	newsIdStr := c.Param(newsId)
	newsId, err := strconv.Atoi(newsIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, `invalid id param, should be int`)
		return
	}

	news, err := h.services.NewsByID(c, newsId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, serverError)
		return
	}

	if news == nil {
		newErrorResponse(c, http.StatusNotFound, `news not found`)
		return
	}

	c.JSON(http.StatusOK, news)
}

func (h *Handler) getNews(c *gin.Context) {
	req := &ReqQuery{}
	if err := c.BindQuery(req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, `invalid query param(s)`)
		return
	}

	news, err := h.services.NewsByFilters(c,
		&newsportal.NewsFilter{
			CategoryID: req.CategoryID,
			TagID:      req.TagID,
		},
		req.PageNum, req.PageSize)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, serverError)
		return
	}

	if len(news) == 0 {
		newErrorResponse(c, http.StatusNotFound, `news not found`)
		return
	}

	res := make([]*News, 0, len(news))
	for _, n := range news {
		res = append(res, newNews(n))
	}

	c.JSON(http.StatusOK, res)
}

func (h *Handler) getNewsCount(c *gin.Context) {
	filter := new(newsportal.NewsFilter)
	if err := c.BindQuery(filter); err != nil {
		newErrorResponse(c, http.StatusBadRequest, `invalid query param(s)`)
		return
	}

	count, err := h.services.NewsCount(c, filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, serverError)
		return
	}

	c.JSON(http.StatusOK, count)
}
