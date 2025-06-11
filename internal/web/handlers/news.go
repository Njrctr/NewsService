package handlers

import (
	"errors"
	"net/http"
	myErrors "news-service/internal/errors"
	"news-service/internal/service"
	"news-service/internal/structs"
	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	newsId = `id`
)

type ReqQuery struct {
	structs.NewsFilter
	PageSize uint `form:"page_size,default=5"`
	PageNum  uint `form:"page_num,default=0"`
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
func getOneNews(c *gin.Context) {
	newsIdStr := c.Param(newsId)
	newsId, err := strconv.Atoi(newsIdStr)
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, `invalid id param, should be int`)
		return
	}

	news, err := service.GetNewsByID(c, newsId)
	if err != nil {
		if errors.Is(err, myErrors.ErrNoRows) {
			newErrorResponse(c, http.StatusNotFound, `news not found`)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, serverError)
		return
	}

	c.JSON(http.StatusOK, news)
}

func getNews(c *gin.Context) {
	req := new(ReqQuery)
	if err := c.BindQuery(req); err != nil {
		newErrorResponse(c, http.StatusBadRequest, `invalid query param(s)`)
		return
	}

	news, err := service.GetNews(c,
		&structs.NewsFilter{
			CategoryID: req.CategoryID,
			TagID:      req.TagID,
		},
		req.PageNum, req.PageSize)
	if err != nil {
		if errors.Is(err, myErrors.ErrNoRows) {
			newErrorResponse(c, http.StatusNotFound, `news not found`)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, serverError)
		return
	}

	c.JSON(http.StatusOK, news)
}

func getNewsCount(c *gin.Context) {
	filter := new(structs.NewsFilter)
	if err := c.BindQuery(filter); err != nil {
		newErrorResponse(c, http.StatusBadRequest, `invalid query param(s)`)
		return
	}

	count, err := service.GetNewsCount(c, filter)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, serverError)
		return
	}

	c.JSON(http.StatusOK, count)
}
