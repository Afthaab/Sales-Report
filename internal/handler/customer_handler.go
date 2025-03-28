package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetAverageValue(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "query is empty",
			"message": "missing start_date or end_date"})
		return
	}

	avgValue, err := h.svc.GetAverageValue(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "could not process the request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   nil,
		"message": "successfully processed the request",
		"response": gin.H{
			"average_value": avgValue,
		},
	})
}

func (h *handler) GetTotalNumberOfOrders(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "query is empty",
			"message": "missing start_date or end_date"})
		return
	}

	totalOrders, err := h.svc.GetTotalOrders(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "could not process the request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   nil,
		"message": "successfully processed the request",
		"response": gin.H{
			"total_orders": totalOrders,
		},
	})
}

func (h *handler) GetTotalCustomers(c *gin.Context) {
	startDate := c.Query("start_date")
	endDate := c.Query("end_date")

	if startDate == "" || endDate == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "query is empty",
			"message": "missing start_date or end_date"})
		return
	}

	totalCustomers, err := h.svc.GetTotalCustomers(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err.Error(),
			"message": "could not process the request",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"error":   nil,
		"message": "successfully processed the request",
		"response": gin.H{
			"total_customers": totalCustomers,
		},
	})

}
