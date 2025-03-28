package script

import (
	"github.com/Afthaab/Sales-Report-Lumel/internal/loader"
	"github.com/Afthaab/Sales-Report-Lumel/internal/util"
)

func RunCSVLoader(loader loader.LoaderInterface) error {
	fileData := loader.LoadCSVFile()

	for filePath, data := range fileData {
		err := loader.StoreTheCSVDateToDb(data)
		if err != nil {
			err := util.MoveFile(filePath, util.ERROR_PATH)
			if err != nil {
				return err
			}
		} else {
			err := util.MoveFile(filePath, util.SUCCESS_PATH)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
