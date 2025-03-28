package repository

import (
	"strconv"
	"time"

	"github.com/Afthaab/Sales-Report-Lumel/internal/model/csvmodel"
	"github.com/Afthaab/Sales-Report-Lumel/internal/model/dbmodel"
	"github.com/Afthaab/Sales-Report-Lumel/internal/util"
	"github.com/rs/zerolog/log"
)

func (r *repo) StoreCustomerData(data csvmodel.Order) error {
	var customerData dbmodel.Customer
	if err := r.db.Where("customer_id = ?", data.CustomerID).First(&customerData).Error; err == nil {
		log.Info().Msg("skipping the data insert; customer id = " + data.CustomerID + " already exists")
	} else if err.Error() == "record not found" {
		customerData = dbmodel.Customer{
			CustomerID:      data.CustomerID,
			CustomerName:    data.CustomerName,
			CustomerEmail:   data.CustomerEmail,
			CustomerAddress: data.CustomerAddr,
		}
		if err := r.db.Create(&customerData).Error; err != nil {
			log.Error().Err(err).Msg("failed to store the data to db; customer_id = " + data.CustomerID)
			return err
		}
	} else {
		log.Error().Err(err).Msg("error occured while checking the exsiting customer; customer_id = " + data.CustomerID)
		return err
	}
	return nil
}

func (r *repo) StoreCategoryData(data csvmodel.Order) (dbmodel.Category, error) {
	var categoryData dbmodel.Category
	if err := r.db.Where("category_name = ?", data.Category).First(&categoryData).Error; err == nil {
		log.Info().Msg("skipping the data insert; category = " + data.Category + " already exists")
	} else if err.Error() == "record not found" {
		categoryData = dbmodel.Category{
			CategoryName: data.Category,
		}
		if err := r.db.Create(&categoryData).Error; err != nil {
			log.Error().Err(err).Msg("failed to store the data to db; category = " + data.Category)
			return dbmodel.Category{}, err
		}
	} else {
		log.Error().Err(err).Msg("error occured while checking the exsiting category; category = " + data.Category)
		return dbmodel.Category{}, err
	}
	return categoryData, r.db.Error
}

func (r *repo) StoreTheRegionData(data csvmodel.Order) (dbmodel.Region, error) {
	var regionData dbmodel.Region
	if err := r.db.Where("region_name = ?", data.Region).First(&regionData).Error; err == nil {
		log.Info().Msg("skipping insert; region = " + data.Region + " already exists")
	} else if err.Error() == "record not found" {
		regionData = dbmodel.Region{
			RegionName: data.Region,
		}
		if err := r.db.Create(&regionData).Error; err != nil {
			log.Error().Err(err).Msg("failed to store the data to db; region = " + data.Region)
			return dbmodel.Region{}, err
		}
	} else {
		log.Error().Err(err).Msg("error occured while checking the exsiting region; region = " + data.Category)
		return dbmodel.Region{}, err
	}
	return regionData, nil
}

func (r *repo) StoreTheProduct(data csvmodel.Order, categoryId int) error {
	var productData dbmodel.Product
	if err := r.db.Where("product_id = ?", data.ProductID).First(&productData).Error; err == nil {
		log.Info().Msg("skipping insert; product_id = " + data.ProductID + " already exists")
	} else if err.Error() == "record not found" {
		productData = dbmodel.Product{
			ProductID:   data.ProductID,
			ProductName: data.ProductName,
			Categoryid:  categoryId,
			UnitPrice:   data.UnitPrice,
		}
		if err := r.db.Create(&productData).Error; err != nil {
			log.Error().Err(err).Msg("failed to store the data to db; product_id = " + data.ProductID)
			return err
		}
	} else {
		log.Error().Err(err).Msg("error occured while checking the exsiting product; product = " + data.ProductID)
		return err
	}
	return nil
}

func (r *repo) StoreTheOrderDetails(data csvmodel.Order, regionId int) error {
	var orderData dbmodel.Order
	if err := r.db.Where("order_id = ?", data.OrderID).First(&orderData).Error; err == nil {
		log.Info().Msg("skipping insert; order_id = " + strconv.Itoa(data.OrderID) + " already exists")
	} else if err.Error() == "record not found" {
		orderDate, err := time.Parse(util.Format, data.DateOfSale)
		if err != nil {
			log.Error().Err(err).Msg("could not parse the date")
			return err
		}
		orderData = dbmodel.Order{
			OrderID:       data.OrderID,
			Customersid:   data.CustomerID,
			OrderDate:     orderDate,
			TotalAmount:   util.TotalAmount(data.QuantitySold, data.UnitPrice, data.ShippingCost, data.Discount),
			ShippingCost:  data.ShippingCost,
			PaymentMethod: data.PaymentMethod,
			Discount:      data.Discount,
			Regionid:      regionId,
		}
		if err := r.db.Create(&orderData).Error; err != nil {
			log.Error().Err(err).Msg("failed to store the data to db; product_id = " + data.ProductID)
			return err
		}
	} else {
		log.Error().Err(err).Msg("error occured while checking the exsiting product; product = " + data.ProductID)
		return err
	}
	return nil
}

func (r *repo) StoreTheOrderItemsDetail(data csvmodel.Order, orderId int, productId string) error {
	orderdItems := dbmodel.OrderItem{
		Ordersid:   orderId,
		Productsid: productId,
		Quantity:   data.QuantitySold,
		UnitPrice:  data.UnitPrice,
	}
	if err := r.db.Create(&orderdItems).Error; err != nil {
		log.Error().Err(err).Msg("failed to store the data to db; product_id = " + data.ProductID)
		return err
	}
	return nil
}
