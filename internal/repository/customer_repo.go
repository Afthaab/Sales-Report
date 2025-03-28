package repository

import (
	"time"

	"github.com/Afthaab/Sales-Report-Lumel/internal/model/dbmodel"
	"github.com/rs/zerolog/log"
)

func (r *repo) GetAverageValue(startdate time.Time, enddate time.Time) (float64, error) {
	var avgOrderValue float64
	var orderDate dbmodel.Order
	err := r.db.Model(&orderDate).
		Where("order_date BETWEEN ? AND ?", startdate, enddate).
		Select("COALESCE(SUM(total_amount) / NULLIF(COUNT(order_id), 0), 0)").
		Scan(&avgOrderValue).Error

	if err != nil {
		log.Error().Err(err).Msg("could not fetch the avg value")
		return 0, err
	}
	return avgOrderValue, err

}

func (r *repo) GetTotalCustomers(startDate, endDate time.Time) (int, error) {
	var totalCustomers int64
	var orderData dbmodel.Order
	err := r.db.Model(&orderData).
		Where("order_date BETWEEN ? AND ?", startDate, endDate).
		Select("COUNT(DISTINCT customersid)").
		Scan(&totalCustomers).Error

	if err != nil {
		log.Error().Err(err).Msg("could not fetch the total customers")
		return 0, err
	}
	return int(totalCustomers), err
}

func (r *repo) GetTotalOrders(startdate time.Time, enddate time.Time) (int, error) {
	var totalOrders int64
	var orderData dbmodel.Order
	err := r.db.Model(&orderData).
		Where("order_date BETWEEN ? AND ?", startdate, enddate).
		Select("COUNT(DISTINCT order_id)").
		Scan(&totalOrders).Error

	if err != nil {
		log.Error().Err(err).Msg("could not fetch the total orders")
		return 0, err
	}
	return int(totalOrders), err
}
