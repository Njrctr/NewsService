package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getTags(c *gin.Context) {
	tags, err := h.services.Tags(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, serverError)
		return
	}

	c.JSON(http.StatusOK, tags)
}
