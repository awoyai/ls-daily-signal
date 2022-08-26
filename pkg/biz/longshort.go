package biz

import (
	"ls-daily-signal/pkg/model"
	"ls-daily-signal/utils"
	"time"
)

type LongShortRatioRepo interface {
	Query(f *model.LongShortFilter) (*model.LongShortRatios, error)
}

type PreOrderRepo interface {
	Query(f *model.PreOrderFilter) ([]*model.PreOrder, error)
}

type SymbolPriceRepo interface {
	Query(f *model.SymbolPriceFilter) (*model.SymbolPrices, error)
}

type LongShortUsecase struct {
	plates    []string
	ratioRepo LongShortRatioRepo
	priceRepo SymbolPriceRepo
	orderRepo PreOrderRepo
}

func (u *LongShortUsecase) GetDailyData(date *time.Time) (map[string][]*model.ExcelData, error) {
	plateMap := make(map[string][]*model.ExcelData)
	for _, plate := range u.plates {
		orders, err := u.orderRepo.Query(&model.PreOrderFilter{CreateDate: date, PlateName: plate})
		if err != nil {
			return nil, err
		}
		var excelDatas []*model.ExcelData
		for _, order := range orders {
			excelData, err := u.getDailyData(date, order)
			if err != nil {
				return nil, err
			}
			excelDatas = append(excelDatas, excelData...)
		}
		plateMap[plate] = excelDatas

	}
	return nil, nil
}

func (u *LongShortUsecase) getDailyData(date *time.Time, order *model.PreOrder) ([]*model.ExcelData, error) {
	data := make([]*model.ExcelData, 2)
	prices, err := u.priceRepo.Query(&model.SymbolPriceFilter{TradingAt: date.Unix(), Symbols: []string{order.Long, order.Short}})
	if err != nil {
		return nil, err
	}
	priceMap := prices.GetMap()
	ratios, err := u.ratioRepo.Query(&model.LongShortFilter{TradingAt: date.Unix(), Varieties: []string{order.Long, order.Short}})
	if err != nil {
		return nil, err
	}
	ratioMap := ratios.GetMap()
	data[0] = u.generateExcelData(priceMap[order.Long].ClosePrice, ratioMap[order.Long].Ratio, order, true)
	data[1] = u.generateExcelData(priceMap[order.Short].ClosePrice, ratioMap[order.Short].Ratio, order, true)
	return data, nil
}

func (u *LongShortUsecase) generateExcelData(price float64, ratio float64, order *model.PreOrder, isLong bool) *model.ExcelData {
	var (
		direction string
		symbol    string
		size      int
	)
	if isLong {
		direction = "多"
		symbol = order.Long
		size = int(order.LongSize)
	} else {
		direction = "空"
		symbol = order.Short
		size = int(order.ShortSize)
	}
	return &model.ExcelData{
		Date:       order.CreateDate.Format(utils.DateLayout),
		Plate:      order.PlateName,
		Direction:  direction,
		Varieties:  utils.ParseVarietyFromSymbol(symbol),
		Symbol:     symbol,
		TodayClose: price,
		Size:       size,
		TodayRatio: ratio,
	}
}
