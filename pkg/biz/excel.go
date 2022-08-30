package biz

import (
	"fmt"
	"ls-daily-signal/pkg/model"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/xuri/excelize/v2"
)

type ExcelUsecase struct {
	plateMap map[string]string
	filePath string
}

func NewExcelUsecase(plateMap map[string]string, filePath string) *ExcelUsecase {
	return &ExcelUsecase{plateMap, filePath}
}

func (u *ExcelUsecase) SaveDaliyDateToExcel(plate string, data []*model.ExcelData) error {
	f, err := excelize.OpenFile(u.filePath)
	if err != nil {
		log.Errorw("SetDaliyDate#OpenFile err", err)
		panic(err)
	}
	defer f.Close()
	sheetName := u.plateMap[plate]
	rows, err := f.GetRows(sheetName)
	if err != nil {
		return err
	}
	index := len(rows)
	start := fmt.Sprintf("A%v", index+1)
	var d  = model.ExcelDatas(data).MergeRepeatSymbol()
	for _, v := range d {
		index++
		axis := fmt.Sprintf("A%v", index)
		params := v.ToSlice(u.plateMap[v.Plate])
		if err := f.SetSheetRow(sheetName, axis, &params); err != nil {
			return err
		}
	}
	end := fmt.Sprintf("A%v", index)
	f.MergeCell(sheetName, start, end)
	if err := f.Save(); err != nil {
		return err
	}
	return nil
}
