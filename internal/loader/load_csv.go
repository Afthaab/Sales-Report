package loader

import (
	"encoding/csv"
	"log"
	"os"
	"path/filepath"

	"github.com/Afthaab/Sales-Report-Lumel/internal/model/csvmodel"
	util "github.com/Afthaab/Sales-Report-Lumel/internal/utils"
	"github.com/jszwec/csvutil"
)

func (l *load) LoadCSVFile() map[string][]csvmodel.Order {
	csvMap := make(map[string][]csvmodel.Order)

	files, err := filepath.Glob(util.CSV_PATH)
	if err != nil {
		log.Printf("Error listing CSV files: %v", err)
		return nil
	}

	for _, filePath := range files {
		log.Printf("Processing file: %s", filePath)

		file, err := os.Open(filePath)
		if err != nil {
			log.Printf("Error opening file %s: %v", filePath, err)
			continue
		}

		reader := csv.NewReader(file)
		headers, err := reader.Read()
		if err != nil {
			log.Printf("Error reading headers in file %s: %v", filePath, err)
			file.Close()
			continue
		}

		decoder, err := csvutil.NewDecoder(reader, headers...)
		if err != nil {
			log.Printf("Error decoding CSV file %s: %v", filePath, err)
			file.Close()
			continue
		}

		var orders []csvmodel.Order
		for {
			var order csvmodel.Order
			if err := decoder.Decode(&order); err != nil {
				break
			}
			orders = append(orders, order)
		}

		file.Close()

		csvMap[filePath] = orders
	}

	return csvMap
}

func (l *load) StoreTheCSVDateToDb(salesdata []csvmodel.Order) error {

	for _, data := range salesdata {

		err := l.repo.StoreCustomerData(data)
		if err != nil {
			return err
		}

		categoryData, err := l.repo.StoreCategoryData(data)
		if err != nil {
			return err
		}

		err = l.repo.StoreTheProduct(data, categoryData.CategoryID)
		if err != nil {
			return err
		}

		regionData, err := l.repo.StoreTheRegionData(data)
		if err != nil {
			return err
		}

		err = l.repo.StoreTheOrderDetails(data, regionData.RegionID)
		if err != nil {
			return err
		}

		err = l.repo.StoreTheOrderItemsDetail(data, data.OrderID, data.ProductID)
		if err != nil {
			return err
		}
	}

	return nil
}
