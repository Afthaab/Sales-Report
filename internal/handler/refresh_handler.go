package handler

import (
	"net/http"

	"github.com/Afthaab/Sales-Report-Lumel/internal/script"
	"github.com/gin-gonic/gin"
)

func (h *handler) RefreshHandler(c *gin.Context) {
	err := script.RunCSVLoader(h.load)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"message": "failed in processing the request",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"error":   nil,
		"message": "successfully processed the request",
	})
}
