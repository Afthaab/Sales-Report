package service

import (
	"time"

	util "github.com/Afthaab/Sales-Report-Lumel/internal/utils"
	"github.com/rs/zerolog/log"
)

func (s *svc) GetAverageValue(startDate string, endDate string) (float64, error) {
	startdate, err := time.Parse(util.Format, startDate)
	if err != nil {
		log.Error().Err(err).Msg("could not parse the date")
		return 0, err
	}
	enddate, err := time.Parse(util.Format, endDate)
	if err != nil {
		log.Error().Err(err).Msg("could not parse the date")
		return 0, err
	}
	avgValue, err := s.repo.GetAverageValue(startdate, enddate)
	if err != nil {
		return 0, err
	}

	return avgValue, nil
}

func (s *svc) GetTotalCustomers(startDate string, endDate string) (int, error) {
	startdate, err := time.Parse(util.Format, startDate)
	if err != nil {
		log.Error().Err(err).Msg("could not parse the date")
		return 0, err
	}
	enddate, err := time.Parse(util.Format, endDate)
	if err != nil {
		log.Error().Err(err).Msg("could not parse the date")
		return 0, err
	}

	totalCustomers, err := s.repo.GetTotalCustomers(startdate, enddate)
	if err != nil {
		return 0, err
	}

	return totalCustomers, nil
}

func (s *svc) GetTotalOrders(startDate string, endDate string) (int, error) {
	startdate, err := time.Parse(util.Format, startDate)
	if err != nil {
		log.Error().Err(err).Msg("could not parse the date")
		return 0, err
	}
	enddate, err := time.Parse(util.Format, endDate)
	if err != nil {
		log.Error().Err(err).Msg("could not parse the date")
		return 0, err
	}

	totalOrders, err := s.repo.GetTotalOrders(startdate, enddate)
	if err != nil {
		return 0, err
	}

	return totalOrders, nil

}
