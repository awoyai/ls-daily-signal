package data

import (
	"ls-daily-signal/pkg/data/xgorm"
	"ls-daily-signal/pkg/model"
)

type symbolPriceRepo struct {
	db Data
}

func NewSymbolPriceRepo(db *xgorm.XGorm) *symbolPriceRepo {
	return &symbolPriceRepo{db: Data{db}}
}

func (r *symbolPriceRepo) Query(f *model.SymbolPriceFilter) (*model.SymbolPrices, error) {
	var res *model.SymbolPrices
	return res, r.db.db.Debug().Where("trading_at = ? and symbol in ?", f.TradingAt, f.Symbols).Find(&res).Error
}
