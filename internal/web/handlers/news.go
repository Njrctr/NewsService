package handlers

import (
	"errors"
	"net/http"
	myErrors "news-service/internal/errors"
	"news-service/internal/structs"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getOneNews(c *gin.Context) {
	news, err := h.news.GetNewsByID(c, 1)
	if err != nil {
		if errors.Is(err, myErrors.ErrNoRows) {
			newErrorResponse(c, http.StatusNotFound, `news not found`)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, news)
}

func (h *Handler) getNews(c *gin.Context) {
	filter := new(structs.NewsFilter)
	news, err := h.news.GetNews(c, filter, 1, 1)
	if err != nil {
		if errors.Is(err, myErrors.ErrNoRows) {
			newErrorResponse(c, http.StatusNotFound, `news not found`)
			return
		}
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, news)
}

func (h *Handler) getNewsCount(c *gin.Context) {

}
