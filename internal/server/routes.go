package server

import (
	"github.com/Afthaab/Sales-Report-Lumel/internal/handler"
	"github.com/gin-gonic/gin"
)

func registerRoutes(router *gin.Engine, handler handler.HandlerInterface) {
	router.POST("/refresh", handler.RefreshHandler)
	router.GET("/total/customers", handler.GetTotalCustomers)
	router.GET("/total/orders", handler.GetTotalNumberOfOrders)
	router.GET("/average/order_value", handler.GetAverageValue)
}
