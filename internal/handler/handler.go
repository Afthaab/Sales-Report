package handler

import (
	"github.com/Afthaab/Sales-Report-Lumel/internal/loader"
	"github.com/Afthaab/Sales-Report-Lumel/internal/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	load loader.LoaderInterface
	svc  service.ServiceInterface
}

type HandlerInterface interface {
	RefreshHandler(c *gin.Context)
	GetTotalCustomers(c *gin.Context)
	GetTotalNumberOfOrders(c *gin.Context)
	GetAverageValue(c *gin.Context)
}

func NewHandlerLayer(load loader.LoaderInterface, svc service.ServiceInterface) HandlerInterface {
	return &handler{
		load: load,
		svc:  svc,
	}
}
