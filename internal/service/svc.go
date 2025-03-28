package service

import "github.com/Afthaab/Sales-Report-Lumel/internal/repository"

type svc struct {
	repo repository.RepoInterface
}

type ServiceInterface interface {
	GetTotalCustomers(startDate string, endDate string) (int, error)
	GetTotalOrders(startDate string, endDate string) (int, error)
	GetAverageValue(startDate string, endDate string) (float64, error)
}

func NewServiceLayer(repo repository.RepoInterface) ServiceInterface {
	return &svc{
		repo: repo,
	}
}
