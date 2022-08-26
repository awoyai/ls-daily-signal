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
	sheetName := u.plateMap[plate]
	lastIndex := f.GetSheetIndex(sheetName)
	var d  = model.ExcelDatas(data).MergeRepeatSymbol()
	for _, v := range d {
		lastIndex++
		axis := fmt.Sprintf("A%v", lastIndex)
		f.SetSheetRow(sheetName, axis, v.ToSlice())
	}
	return nil
}
