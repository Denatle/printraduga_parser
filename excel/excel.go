package excel

import (
	"errors"
	"log"
	"printraduga_parser/shared"

	"github.com/xuri/excelize/v2"
)

type DefaultExcelWriter struct {
}

func (d DefaultExcelWriter) Write(filename string, data []shared.CostInfo) error {
	for _, cost := range data {
		log.Println(cost)
	}
	f := excelize.NewFile()
	defer func() error {
		if err := f.Close(); err != nil {
			return err
		}
		return nil
	}()
	if _, err := f.NewSheet("Прозрачные"); err != nil {
		return err
	}
	if _, err := f.NewSheet("Голографические"); err != nil {
		return err
	}

	f.DeleteSheet("Sheet1")

	for idx, costinfo := range data {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			return err
		}
		var sheetName string
		switch costinfo.ParserType {
		case shared.Translusent:
			sheetName = "Прозрачные"
		case shared.Holo:
			sheetName = "Голографические"
		default:
			return errors.New("Invalid CostInfo: ParserType not found")
		}

		log.Println(sheetName)
		if err := f.SetSheetRow(sheetName, cell, &[]interface{}{
			costinfo.Name, costinfo.Link, costinfo.Cost,
		}); err != nil {
			return err
		}
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	return nil
}
