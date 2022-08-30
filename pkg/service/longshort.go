package service

import (
	"ls-daily-signal/pkg/biz"
	"time"

	"github.com/go-kratos/kratos/v2/log"
)

type LongShrotService struct {
	excelUc *biz.ExcelUsecase
	lsUc    *biz.LongShortUsecase
}

func NewExcelUsecase(excelUc *biz.ExcelUsecase, lsUc *biz.LongShortUsecase) *LongShrotService {
	return &LongShrotService{excelUc: excelUc, lsUc: lsUc}
}

func (s *LongShrotService) CreateDailySignal(date *time.Time) error {
	dataMap, err := s.lsUc.GetDailyData(date)
	if err != nil {
		return err
	}
	for k, v := range dataMap {
		plateName, data := k, v
		if err := s.excelUc.SaveDaliyDateToExcel(plateName, data); err != nil {
			log.Errorw("CreateDailySignal#SaveDaliyDateToExcel err", err)
		}

	}
	return nil
}
