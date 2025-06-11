package handlers

import (
	"net/http"
	"news-service/internal/service"

	"github.com/gin-gonic/gin"
)

func getCategories(c *gin.Context) {
	cats, err := service.GetCategories(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, serverError)
		return
	}

	c.JSON(http.StatusOK, cats)
}
