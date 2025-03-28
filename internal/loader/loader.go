package loader

import (
	"github.com/Afthaab/Sales-Report-Lumel/internal/model/csvmodel"
	"github.com/Afthaab/Sales-Report-Lumel/internal/repository"
	"gorm.io/gorm"
)

type load struct {
	db   *gorm.DB
	repo repository.RepoInterface
}

type LoaderInterface interface {
	LoadCSVFile() map[string][]csvmodel.Order
	StoreTheCSVDateToDb(salesdata []csvmodel.Order) error
}

func NewLoader(db *gorm.DB, repo repository.RepoInterface) LoaderInterface {
	return &load{
		db:   db,
		repo: repo,
	}
}
