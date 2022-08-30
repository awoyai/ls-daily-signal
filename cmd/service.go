package main

import (
	"encoding/json"
	"ls-daily-signal/pkg/biz"
	"ls-daily-signal/pkg/data"
	"ls-daily-signal/pkg/model"
	"ls-daily-signal/pkg/service"
)

func initService(conf *model.Conf) (*service.LongShrotService, func(), error) {
	db, close, err := data.NewXGorm(&conf.Data)
	if err != nil {
		panic(err)
	}
	longShortRatioRepo := data.NewLongShortRatioRepo(db)
	preOrderRepo := data.NewPreOrderRepo(db)
	symbolPriceRepo := data.NewSymbolPriceRepo(db)

	var (
		plateName map[string]string
		plates    []string
	)
	if err := json.Unmarshal(conf.PlateConfig.PlateName, &plateName); err != nil {
		close()
		return nil, nil, err
	}

	for plate := range plateName {
		plates = append(plates, plate)
	}
	excelUsecase := biz.NewExcelUsecase(plateName, conf.PlateConfig.FilePath)
	longshortUsecase := biz.NewLongShortUsecase(plates, longShortRatioRepo, symbolPriceRepo, preOrderRepo)

	return service.NewExcelUsecase(excelUsecase, longshortUsecase), close, nil
}
