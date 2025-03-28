package repository

import (
	"time"

	"github.com/Afthaab/Sales-Report-Lumel/internal/model/csvmodel"
	"github.com/Afthaab/Sales-Report-Lumel/internal/model/dbmodel"
	"gorm.io/gorm"
)

type repo struct {
	db *gorm.DB
}

type RepoInterface interface {
	StoreCustomerData(data csvmodel.Order) error
	StoreCategoryData(data csvmodel.Order) (dbmodel.Category, error)
	StoreTheRegionData(data csvmodel.Order) (dbmodel.Region, error)
	StoreTheProduct(data csvmodel.Order, categoryId int) error
	StoreTheOrderDetails(data csvmodel.Order, regionId int) error
	StoreTheOrderItemsDetail(data csvmodel.Order, orderId int, productId string) error
	GetTotalCustomers(startDate, enddate time.Time) (int, error)
	GetTotalOrders(startdate time.Time, enddate time.Time) (int, error)
	GetAverageValue(startdate time.Time, enddate time.Time) (float64, error)
}

func NewRepoLayer(db *gorm.DB) RepoInterface {
	return &repo{
		db: db,
	}
}
