package util

import (
	"log"
	"os"
	"path/filepath"
)

func TotalAmount(QuantitySold int, UnitPrice float64, ShippingCost float64, discount float64) float64 {
	return float64((QuantitySold * int(UnitPrice)) - int(discount) + int(ShippingCost))

}

func MoveFile(filePath, targetDir string) error {
	fileName := filepath.Base(filePath)
	newPath := filepath.Join(targetDir, fileName)

	err := os.Rename(filePath, newPath)
	if err != nil {
		log.Printf("error moving file %s to %s: %v", filePath, targetDir, err)
	} else {
		log.Printf("moved file %s to %s", filePath, targetDir)
	}
	return err
}
