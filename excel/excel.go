package excel

import (
	"printraduga_parser/shared"

	"github.com/xuri/excelize/v2"
)

type DefaultExcelWriter struct {
}

func (d DefaultExcelWriter) Write(filename string, data map[string][]shared.CostData) error {
	f := excelize.NewFile()
	defer func() error {
		if err := f.Close(); err != nil {
			return err
		}
		return nil
	}()

	for sheet, costs := range data {
		if _, err := f.NewSheet(sheet); err != nil {
			return err
		}
		for i, cost := range costs {
			cell, err := excelize.CoordinatesToCellName(1, i+1)
			if err != nil {
				return err
			}

			if err := f.SetSheetRow(sheet, cell, &[]interface{}{
				cost.Name, cost.Link, cost.Cost,
			}); err != nil {
				return err
			}
		}

	}

	f.DeleteSheet("Sheet1")
	// Save spreadsheet by the given path.
	if err := f.SaveAs(filename); err != nil {
		return err
	}
	return nil
}
