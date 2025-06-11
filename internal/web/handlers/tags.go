package handlers

import (
	"net/http"
	"news-service/internal/service"

	"github.com/gin-gonic/gin"
)

func getTags(c *gin.Context) {
	tags, err := service.GetTags(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, serverError)
		return
	}

	c.JSON(http.StatusOK, tags)
}
