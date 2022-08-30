package data

import (
	"ls-daily-signal/pkg/data/xgorm"
	"ls-daily-signal/pkg/model"
)

type longShortRatioRepo struct {
	db Data
}

func NewLongShortRatioRepo(db *xgorm.XGorm) *longShortRatioRepo {
	return &longShortRatioRepo{db: Data{db}}
}

func (r *longShortRatioRepo) Query(f *model.LongShortFilter) (*model.LongShortRatios, error) {
	var res *model.LongShortRatios
	return res, r.db.db.Debug().Where("trading_at = ? and variety in ?", f.TradingAt, f.Varieties).Find(&res).Error
}
