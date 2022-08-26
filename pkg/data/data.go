package data

import (
	"ls-daily-signal/pkg/data/xgorm"
	"ls-daily-signal/pkg/model"
)

type Data struct {
	db *xgorm.XGorm
}

func NewXGorm(c *model.Data) (*xgorm.XGorm, func(), error) {
	return xgorm.OpenDB(&xgorm.Config{
		DBType:          xgorm.DBType(c.Driver),
		DBUser:          c.DbUser,
		DBPassword:      c.DbPassword,
		DBUrl:           c.DbUrl,
		Database:        c.DbName,
		DryRun:          c.DryRun,
		MaxIdleConns:    int(c.MaxIdleConns),
		MaxOpenConns:    int(c.MaxOpenConns),
		ConnMaxLifeTime: c.ConnMaxLifeTime,
	})
}
