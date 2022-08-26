package data

import (
	"ls-daily-signal/pkg/data/xgorm"
	"ls-daily-signal/pkg/model"
)

type symbolPriceRepo struct {
	db Data
}

func NewsymbolPriceRepo(db *xgorm.XGorm) *symbolPriceRepo {
	return &symbolPriceRepo{db: Data{db}}
}

func (r *symbolPriceRepo) Query(f *model.SymbolPriceFilter) (*model.SymbolPrices, error) {
	var res *model.SymbolPrices
	return res, r.db.db.Where("trading_at = date(?) and symbol in ?", f.TradingAt, f.Symbols).Find(&res).Error
}
