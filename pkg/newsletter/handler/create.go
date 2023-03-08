package handler

import (
	"net/http"

	"git.mcontigo.com/safeplay/newsletter-api/pkg/newsletter"
	"github.com/gin-gonic/gin"
)

func (h *handler) Create(c *gin.Context) {
	params := &newsletter.Subscription{}

	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.svc.Create(c, params.UserID, params.BlogID, params.Interests)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "OK")
}
