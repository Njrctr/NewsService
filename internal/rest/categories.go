package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getCategories(c *gin.Context) {
	cats, err := h.services.GetCategories(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, serverError)
		return
	}

	c.JSON(http.StatusOK, cats)
}
