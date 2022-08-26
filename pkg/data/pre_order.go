package data

import (
	"ls-daily-signal/pkg/data/xgorm"
	"ls-daily-signal/pkg/model"
)

type preOrderRepo struct {
	db Data
}

func NewpreOrderRepo(db *xgorm.XGorm) *preOrderRepo {
	return &preOrderRepo{db: Data{db}}
}

func (r *preOrderRepo) Query(f *model.PreOrderFilter) ([]*model.PreOrder, error) {
	var res []*model.PreOrder
	return res, r.db.db.Where("create_date = date(?) and plate_name = ?", f.CreateDate, f.PlateName).Find(&res).Error
}
